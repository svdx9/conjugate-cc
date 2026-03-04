# claude (repo root)

## worktree discipline (mandatory)

NEVER write, edit, or delete code files on the current branch directly.

Before touching any code:
1. Create a new git worktree: use the EnterWorktree tool (Claude Code) or `git worktree add` (CLI)
2. All implementation work happens inside that worktree on its own branch
3. When done, open a PR from the worktree branch into main

This applies to every task, no exceptions. Planning, reading, and backlog updates do not require a worktree.

## interaction protocol (mandatory)

when information is missing, ambiguous, or there are multiple plausible implementations:
- ask exactly one question at a time.
- questions must be yes/no only.
- do not ask multi-part questions.
- do not offer options lists inside the question.
- after I answer, either:
  - proceed with implementation, or
  - ask the next single yes/no question.

- if a yes/no question is genuinely impossible:
  - ask a yes/no question that selects the default path, and state the default explicitly.
    example: "no = use default X, yes = you will provide Y"

- do not continue coding while waiting for an answer when a blocking question is open.
- if no clarification is needed, do not ask questions.


rules:
- prefer minimal diffs.
- if working under web/*, obey web/CLAUDE.md in addition to this file.
- if rules conflict, prefer the more specific (deeper path) CLAUDE.md.
- never edit files outside the current worktree, even to unblock other work.

## mandatory workflow steps

- after editing any Go file, run `make format`.
- before committing, run `make test`.

## worktree naming (mandatory)

Always pass a descriptive `name` to EnterWorktree — never use the random default.

Format: `task-<id>-<short-slug>` (kebab-case, matches the backlog task)

Examples:
- `task-4-add-ci` → worktree `.claude/worktrees/task-4-add-ci`, branch `claude/task-4-add-ci`
- `task-1-scaffold` → worktree `.claude/worktrees/task-1-scaffold`, branch `claude/task-1-scaffold`

If the work has no backlog task, use a short descriptive slug (e.g. `fix-air-config`).

# skills

- required skill:
  - security best practices: @.claude/skills/security-best-practices/SKILL.md
- trigger rule:
  - when a task changes authentication, authorization, password handling, session/cookie handling, input validation, injection prevention, or security logging, load and follow the security skill before coding.

## golang backend skill (mandatory)

all golang backend implementation rules are defined in:

@.claude/skills/go-backend/SKILL.md

this skill is an enforcement contract, not guidance.

## automatic backend detection rule (mandatory)

a task requires loading the golang backend skill if it:

- modifies files under backend/
- changes http handlers, middleware, routing, or server configuration
- changes services, domain logic, or orchestration
- changes store/database logic
- changes configuration loading
- introduces or modifies migrations
- affects api contracts or oapi-codegen output
- modifies tests for backend packages

if any of the above are true, the agent MUST load:

@.claude/skills/go-backend/SKILL.md

the agent must explicitly state that the golang backend skill was loaded before implementation.


<!-- BACKLOG.MD MCP GUIDELINES START -->

<CRITICAL_INSTRUCTION>

## BACKLOG WORKFLOW INSTRUCTIONS

This project uses Backlog.md MCP for all task and project management activities.

**CRITICAL GUIDANCE**

- If your client supports MCP resources, read `backlog://workflow/overview` to understand when and how to use Backlog for this project.
- If your client only supports tools or the above request fails, call `backlog.get_workflow_overview()` tool to load the tool-oriented overview (it lists the matching guide tools).

- **First time working here?** Read the overview resource IMMEDIATELY to learn the workflow
- **Already familiar?** You should have the overview cached ("## Backlog.md Overview (MCP)")
- **When to read it**: BEFORE creating tasks, or when you're unsure whether to track work

These guides cover:
- Decision framework for when to create tasks
- Search-first workflow to avoid duplicates
- Links to detailed guides for task creation, execution, and finalization
- MCP tools reference

You MUST read the overview resource to understand the complete workflow. The information is NOT summarized here.

</CRITICAL_INSTRUCTION>

<!-- BACKLOG.MD MCP GUIDELINES END -->
