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

    const { data: mergedPRs } = await octokit.pulls.list({
      owner: 'AvineshTripathi',
      repo: 'cred',
      state: 'closed',
      sort: 'updated',
      direction: 'desc',
    });

    mergedPRs.forEach((pr) => {
      const { number, user, labels } = pr;
      const author = user.login;

      if (!authorData[author]) {
        authorData[author] = {
          link: `link_to_${author}_profile`,
        };
      }

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
    });

    console.log(authorData)
    console.log(process.env.GITHUB_WORKSPACE+"/analyze/author.json")
    fs.writeFileSync("./author.json", JSON.stringify(authorData, null, 2));
    
    console.log('PRs processed and author.json updated successfully.');
  } catch (error) {
    console.error('Error occurred:', error);
  }

  const f = fs.readFileSync("./author.json", 'utf8');
  console.log(f)
}


async function commitChangesAndCreatePR() {
  try {
    const owner = 'AvineshTripathi';
    const repo = 'cred';
    const branchName = 'new';

    // Read the updated/created file content
    const filePath = './author.json'; // Adjust as per your file path
    const fileContent = fs.readFileSync(filePath, 'utf-8');

    // Create or update the file in the repository
    await octokit.repos.createOrUpdateFileContents({
      owner,
      repo,
      path: filePath,
      branch: branchName,
      message: 'Updated author.json', // Commit message
      content: Buffer.from(fileContent).toString('base64'),
    });

    // Create a pull request
    const { data: pullRequest } = await octokit.pulls.create({
      owner,
      repo,
      title: 'Update author.json',
      head: branchName,
      base: 'main', // Change the base branch as needed
      body: 'Changes to author.json', // PR description
    });

    console.log('Pull request created:', pullRequest.html_url);
  } catch (error) {
    console.error('Error occurred:', error);
  }
}

async function main() {
  try {
    await processPRs(); // Wait for processPRs to complete
    await commitChangesAndCreatePR(); // Wait for commitAndCreatePR to complete
  } catch (error) {
    console.error('Error occurred in main:', error);
  }
}

main();
