import {
  brightGreen,
  brightRed,
} from "https://deno.land/std@0.160.0/fmt/colors.ts";
import { config } from "https://deno.land/x/dotenv@v3.2.0/mod.ts";
const ENV = config();

function* range(min: number, max: number) {
  let i = min;
  while (i < max) {
    yield i++;
  }
}

function error(message: string) {
  console.error(brightRed("Error:"), message);
  Deno.exit(1);
}

function isNumber(n: any) {
  return typeof n == "number" && !isNaN(n - n);
}

let year: number;

async function main() {
  const SESSION_ID = ENV.AOC_SESSION_ID;

  if (!SESSION_ID) {
    error("Missing: AOC_SESSION_ID field from .env");
  }

  if (Deno.args.length == 0) {
    error("Missing `year` argument. Please specify the year");
  }

  year = parseInt(Deno.args[0]);

  if (!isNumber(year)) {
    error(
      "Invalid `year` argument. Please enter atleast one valid year",
    );
  }

  for (let i = 0; i < Deno.args.length; i++) {
    const y = parseInt(Deno.args[i]);
    if (isNumber(y)) {
      year = y;

      for (const day of range(1, 26)) {
        const success = await generateScaffolding(day);
        if (!success) {
          break;
        }
      }
      console.log("\n");
    }
  }
}

async function generateScaffolding(day: number) {
  const folder = `./${year}/${day}`;
  const filename = `${folder}/input.txt`;
  const headers = new Headers();
  headers.set(
    "Cookie",
    `session=${ENV.AOC_SESSION_ID}`,
  );
  const response = await fetch(
    `https://adventofcode.com/${year}/day/${day}/input`,
    {
      headers,
    },
  );

  if (!response.ok) {
    return false;
  }

  // create directory and file if they both dont exist
  // attempt to create folder

  // Validate year folder
  try {
    await Deno.lstat(`./${year}`);
  } catch (error) {
    Deno.mkdirSync(`./${year}`);
  }

  // Validate each day
  try {
    await Deno.lstat(folder);
  } catch (error) {
    Deno.mkdirSync(folder);
  }

  // attempt to open the file and if it exists then return true
  try {
    await Deno.lstat(filename);
    return true;
  } catch (error) {}

  const file = await Deno.create(filename);
  await response.body?.pipeTo(file.writable);
  const DEFAULT_FILE_CONTENTS = `

  const FILE_NAME = "./input.txt";
  const input = await Deno.readTextFile(FILE_NAME);

  export async function step1 () {
    // handle first step here ...
  }

  export async function step2 () {
    // handle second step here ...
  } 

  `;

  Deno.writeTextFile(folder + "/main.ts", DEFAULT_FILE_CONTENTS);
  console.log(`${brightGreen("+")} ${year}/${day}`);
  return true;
}

try {
  main();
} catch (err) {
  error(err);
}
