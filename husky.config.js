export default {
  hooks: {
    'pre-commit': 'cd frontend && npm run typecheck && npm run lint && npm run test'
  }
}
