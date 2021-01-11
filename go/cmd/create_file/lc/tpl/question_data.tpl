{
  "operationName": "questionData",
  "variables": { "titleSlug": "{{.titleSlug}}" },
  "query": "query questionData($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionFrontendId\n    title\n    content\n    difficulty\n    codeSnippets {\n      langSlug\n      code\n}\n}\n}\n"
}
