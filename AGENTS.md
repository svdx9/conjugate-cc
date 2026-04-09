
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

---

## SKILLS

skills are located in .claude/skills

## TECHNOLOGY STACK CONTEXT (for agents)

**Frontend:**
- Framework: SolidJS 1.9.3 (NOT React)
- Language: TypeScript
- Styling: Tailwind CSS 4.0.0
- Build Tool: Vite 6.0.7
- State Management: SolidJS signals (`createSignal()`)
- Testing: Vitest + @solidjs/testing-library
- API Client: openapi-fetch

**Backend:**
- Language: Go 1.25.7
- HTTP Server: Standard library net/http
- API: OpenAPI-based REST endpoints
- Testing: Standard library testing package

**Key Implementation Details:**
- Frontend components use `.tsx` extension
- Backend uses Go modules (go.mod)
- API specification in `docs/schema/v1/api.yaml`
- Frontend API client generated from OpenAPI spec
- State management via SolidJS reactivity primitives

**Testing Approach:**
- Frontend: Vitest unit + component tests
- Backend: Go standard library tests
- Integration: Manual + curl verification

Last updated: 2026-04-09
