-- name: CreateOrUpdateMagicLinkToken :one
INSERT INTO magic_links (user_id, token_hash, expires_at, consumed_at)
VALUES ($1, $2, $3, NULL)
ON CONFLICT (user_id) WHERE consumed_at IS NULL
DO UPDATE SET
	token_hash = EXCLUDED.token_hash,
	created_at = CURRENT_TIMESTAMP,
	expires_at = EXCLUDED.expires_at,
	consumed_at = NULL
RETURNING id, user_id, token_hash, expires_at, consumed_at, created_at;

-- name: FindMagicLinkByTokenHash :one
SELECT
  ml.id,
  ml.user_id,
  ml.token_hash,
  ml.expires_at,
  ml.consumed_at,
  ml.created_at,
  u.email
FROM magic_links ml
JOIN users u ON u.id = ml.user_id
WHERE ml.token_hash = $1
  AND ml.consumed_at IS NULL
  AND ml.expires_at > now();

-- name: ConsumeMagicLink :exec
UPDATE magic_links
SET consumed_at = now()
WHERE id = $1 AND consumed_at IS NULL;
