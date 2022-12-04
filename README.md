# AOC
 Advent Of Code challenges. This repository contains all the code written for my AOC challenges. It Also contains a helpful cli script to generate each folder and input file. Simply add a `.env` file and place your session in there and this will pull down the daily inputs.


## Using This Repo

If you could like to use this repo and the cli tool and have Deno installed then simply run 
```bash
deno run -A aoc.ts 2022 2016
```

This will generate the folder structure for the years 2016 and 2022. As long as you pass one year then it will generate the scaffolding for that year. Doing this allows you to very quickly generate any year you would like. If a say is not yet released for a year then the generation will only happen for days that are already live.