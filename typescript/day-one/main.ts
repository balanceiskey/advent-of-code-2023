import { readAndSplit } from "../utils/readAndSplit.ts";

export function getScore(input: string): number {
  const matches = input.match(/\d/g);
  if (!matches || !matches.length) {
    throw new Error("you passed a string without digits");
  }
  const joined = `${matches[0]}${matches[matches.length - 1]}}`;
  return parseInt(joined);
}

const STRING_TO_VALUE: Record<string, number> = {
  one: 1,
  two: 2,
  three: 3,
  four: 4,
  five: 5,
  six: 6,
  seven: 7,
  eight: 8,
  nine: 9,
};

export function getScoreCorrected(input: string): number {
  const words = Object.keys(STRING_TO_VALUE);

  // find matching words with positions first.
  const matches: RegExpMatchArray[] = [];

  words.forEach((w) => {
    const re = new RegExp(w, "g");
    const wordMatches = input.matchAll(re);

    for (const match of wordMatches) {
      matches.push(match);
    }
  });

  const matchesNumeric = input.matchAll(/\d/g);

  for (const match of matchesNumeric) {
    matches.push(match);
  }

  const sorted = matches.sort((a: RegExpMatchArray, b: RegExpMatchArray) => {
    if (
      typeof a.index !== "number" ||
      typeof b.index !== "number" ||
      a.index < 0 ||
      b.index < 0
    ) {
      throw new Error("index missing from match");
    }

    if (a.index < b.index) {
      return -1;
    } else if (a.index > b.index) {
      return 1;
    }
    return 0;
  });

  const firstValue = sorted[0]
  const secondValue = sorted[sorted.length - 1]

  const firstDigit = STRING_TO_VALUE[firstValue[0]] ?? firstValue[0]
  const secondDigit = STRING_TO_VALUE[secondValue[0]] ?? secondValue[0]

  return parseInt(`${firstDigit}${secondDigit}`)
}

// pt 1

const lines = await readAndSplit('1.1.txt')

const sum = lines.map(l => getScore(l)).reduce((acc, val) => acc + val, 0)

console.log(`Sum of part one: ${sum}`)

// pt 2

const lines2 = await readAndSplit('1.2.txt')

const sum2 = lines.map(l => getScoreCorrected(l)).reduce((acc, val) => acc + val, 0)
console.log(`Sum of part two: ${sum2}`)
