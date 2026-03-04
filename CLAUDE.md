# claude (repo root)

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

## mandatory workflow steps

- after editing any Go file, run `make format`.
- before committing, run `make test`.

# skills

- required skill:
  - security best practices: @.agents/skills/security-best-practices/SKILL.md
- trigger rule:
  - when a task changes authentication, authorization, password handling, session/cookie handling, input validation, injection prevention, or security logging, load and follow the security skill before coding.

## golang backend skill (mandatory)

all golang backend implementation rules are defined in:

@.agents/skills/go-backend/SKILL.md

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

@.agents/skills/go-backend/SKILL.md

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
