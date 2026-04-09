// Shared TypeScript types and utilities for the application

/**
 * Standard Result type for operations that may fail.
 * Uses a Discriminated Union (ok: true | false) to ensure that the
 * success path (data) and error path never overlap.
 * 
 * This pattern MUST be used for all data operations, API calls, and
 * any function that can fail in a non-exceptional way.
 */
export type Result<T> =
  | { 
      ok: true; 
      data: T; 
    }
  | { 
      ok: false; 
      error: string; 
      code: string;
      details?: Record<string, unknown>;
    };

/**
 * Type guard for error checking.
 * Usage: if (isError(result)) { handle error }
 */
export function isError<T>(result: Result<T>): result is Extract<Result<T>, { ok: false }> {
  return !result.ok;
}

/**
 * Type guard for success checking.
 * Usage: if (isSuccess(result)) { use result.data }
 */
export function isSuccess<T>(result: Result<T>): result is Extract<Result<T>, { ok: true }> {
  return result.ok;
}

/**
 * Common error codes for consistent error handling across the application.
 * Feature-specific modules MAY extend this with domain-specific codes.
 */
export type CommonErrorCode =
  | 'INVALID_INPUT'
  | 'NOT_FOUND'
  | 'NETWORK_ERROR'
  | 'UNAUTHORIZED'
  | 'FORBIDDEN'
  | 'UNKNOWN';

/**
 * Utility function to create success results.
 * Usage: return success(data)
 */
export function success<T>(data: T): Result<T> {
  return { ok: true, data };
}

/**
 * Utility function to create error results.
 * Usage: return error('Something went wrong', 'NOT_FOUND')
 */
export function error<T>(error: string, code: string, details?: Record<string, unknown>): Result<T> {
  return { ok: false, error, code, details };
}

