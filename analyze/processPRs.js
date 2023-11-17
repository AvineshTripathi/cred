const { Octokit } = require('@octokit/rest');
const fs = require('fs');
const path = require('path');

const githubToken = process.env.GH_TOKEN;
const octokit = new Octokit({
  auth: githubToken,
});

async function processPRs() {
  try {
    let authorData = {};

    console.log("workspace", process.env.GITHUB_WORKSPACE)
    const authorJsonPath = path.join(process.env.GITHUB_WORKSPACE, 'author.json');

    if (!fs.existsSync(authorJsonPath)) {
      console.log("existsSync", fs.existsSync(authorJsonPath))
      fs.writeFileSync(authorJsonPath, JSON.stringify(authorData, null, 2));
      console.log('author.json created.');
    } else {
      const fileContent = fs.readFileSync(authorJsonPath, 'utf8');
      if (fileContent.trim() !== '') {
        authorData = JSON.parse(fileContent);
      } else {
        console.log('author.json is empty or contains only whitespace.');
      }
    }

    const currentDate = new Date();
    const sevenDaysAgo = new Date(currentDate);
    sevenDaysAgo.setDate(sevenDaysAgo.getDate() - 1);

    const { data: mergedPRs } = await octokit.pulls.list({
      owner: 'AvineshTripathi',
      repo: 'cred',
      state: 'closed',
      sort: 'updated',
      direction: 'desc',
    });

    const recentPRs = mergedPRs.filter(pr => {
      const prCreatedAt = new Date(pr.created_at);
      return prCreatedAt > sevenDaysAgo;
    });

    recentPRs.forEach((pr) => {
      const { number, user, labels } = pr;
      const author = user.login;

      if (!authorData[author]) {
        authorData[author] = {
          link: `link_to_${author}_profile`,
        };
      }

      if (labels.length > 0) {
        labels.forEach((label) => {
          const labelName = label.name;

          if (!authorData[author][labelName]) {
            authorData[author][labelName] = [number];
          } else {
            if (!authorData[author][labelName].includes(number)) {
              authorData[author][labelName].push(number);
            }
          }
        });
      } else {
        console.log(`Skipping PR #${number}: No labels found`);
      }
    });

    fs.writeFileSync(authorJsonPath, JSON.stringify(authorData, null, 2));
    console.log('PRs processed and author.json updated successfully.');
  } catch (error) {
    console.error('Error occurred:', error);
  }
}

processPRs();
  
