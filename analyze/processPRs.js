const { Octokit } = require('@octokit/rest');
const fs = require('fs');

const githubToken = process.env.GH_TOKEN;
// Initialize Octokit with your GitHub token
const octokit = new Octokit({
  auth: githubToken,
});

async function processPRs() {
  try {
    let authorData = {};

    // Read the existing author data from author.json
    if (fs.existsSync('author.json')) {
      const fileContent = fs.readFileSync('author.json', 'utf8');
      if (fileContent.trim() !== '') {
        authorData = JSON.parse(fileContent);
      } else {
        console.log('author.json is empty or contains only whitespace.');
      }
    } else {
      console.log('author.json does not exist.');
    }

    console.log('Existing author data:', authorData);

    // Fetch merged PRs in the last week
    const { data: mergedPRs } = await octokit.pulls.list({
      owner: 'AvineshTripathi',
      repo: 'cred',
      state: 'closed',
      sort: 'updated',
      direction: 'desc',
    });

    // Process each merged PR
    mergedPRs.forEach((pr) => {
      const { number, user, labels } = pr;
      const author = user.login;

      // If author doesn't exist, create a new entry
      if (!authorData[author]) {
        authorData[author] = {
          link: `link_to_${author}_profile`,
        };
      }

      // Process labels for this PR
      labels.forEach((label) => {
        const labelName = label.name;

        // Check if the label exists for the author in author.json
        if (!authorData[author][labelName]) {
          // If the label doesn't exist for the author, create a new array and add PR number
          authorData[author][labelName] = [number];
        } else {
          // If the label exists, push PR number to the corresponding label array
          authorData[author][labelName].push(number);
        }
      });
    });

    // Update author.json with the modified data
    fs.writeFileSync('author.json', JSON.stringify(authorData, null, 2));
    
    console.log('PRs processed and author.json updated successfully.');
  } catch (error) {
    console.error('Error occurred:', error);
  }
}

processPRs();
