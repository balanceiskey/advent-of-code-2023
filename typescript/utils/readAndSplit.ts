export async function readAndSplit(fileName: string) {
  const resolvedPath = `../../puzzle-inputs/${fileName}`
  const text = await Deno.readTextFileSync(resolvedPath)  
  return text.split("\n")
}