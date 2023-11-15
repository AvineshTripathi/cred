# Cred

This tool is in initial developement phase and is build with the purpose to scan the git repo for AWS creds.


## How to run?

Currently in order to run the codebase one has to clone the repo with 
```bash
git clone https://github.com/AvineshTripathi/cred.git
```


In order to run it 

```bash
go run .
``` 

The program will ask you for a repo link and that will be the repo the code will scan.

## Approach

Initial idea was to use the go-github lib to get a content iterator and search for the regexp match of AWS creds(both the components).
However the problem faced was in the iteration. A work around found was to clone the repo in tmp dir and iterate it.


## Architecture 

Programs starts from `main.go` file which first asks for scan selections for the user. Currently there are two types of scan namely Repo scan(`readRepo.go`) and Commit scan(`readCommit.go`). Both the scans have their respective functions and each functions returns 2 slices i.e. possible access id slice and possible access secret key slice.

These slices are sent to the `validate.go` where all the possible combinations of the pairs are made and these pairs are used to call the AWS api using AWS-v2-go. The response of the api call is the final check for the pair if they are actual creds present in the repo. 


## Problems to be tackled 

Currently the sub-string are found using the regexp module however there are some edge cases where the actual creds can be present in different form and therefore can go undetected. 

Second major concern is the speed, when some static files or ingeneral certain files like `package-lock.json` are present in the repo the possible secret key found in the repo by the detector is too high and therefore the performing tests for every pair ends up taking a lot time(even though goroutine is used)
- One workaround found was to skip those files(mostly static files generate by the code) as those have a very less chances to have creds present however when the tool when exposed to variety of repo may encounter different repos that have varieties of such file so detection of those file can be challenging.


## What's Next

Apart from finding a concret solution for the above mentioned problems, the tool needs a better user interface hence maybe adding cobra cli to this is better to give it a new look.

To make the detection more accurate research should be done. At the moment one of the improvements that might help the tool is to add some AI support that detects the creds(not sure if that could be added at this stage of development).

There is a manual work seen here right now where a user has to manually use this tool to check for the creds, possible improvement here is to add this tool as a part of the developement environment 
