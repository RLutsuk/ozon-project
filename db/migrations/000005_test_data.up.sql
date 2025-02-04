INSERT INTO users (id, username, email, first_name, last_name) VALUES
  ('11111111-1111-1111-1111-111111111111', 'alice', 'alice@example.com', 'Alice', 'Johnson'),
  ('22222222-2222-2222-2222-222222222222', 'bob', 'bob@example.com', 'Bob', 'Smith'),
  ('33333333-3333-3333-3333-333333333333', 'charlie', 'charlie@example.com', 'Charlie', 'Brown');

  INSERT INTO posts (id, title, body, user_id, allow_comments) VALUES
  ('aaaaaaa1-aaaa-aaaa-aaaa-aaaaaaaaaaa1', 'Hello World', 'This is my first post on this platform!', '11111111-1111-1111-1111-111111111111', TRUE),
  ('aaaaaaa2-aaaa-aaaa-aaaa-aaaaaaaaaaa2', 'Test Post', 'This is my second post!', '11111111-1111-1111-1111-111111111111', FALSE),
  ('aaaaaaa3-aaaa-aaaa-aaaa-aaaaaaaaaaa3', 'Tech Trends 2025', 'A deep dive into upcoming technologies.', '22222222-2222-2222-2222-222222222222', TRUE),
  ('aaaaaaa4-aaaa-aaaa-aaaa-aaaaaaaaaaa4', 'No Comments Allowed', 'This post is for reading only.', '33333333-3333-3333-3333-333333333333', FALSE);

  INSERT INTO comments (id, body, user_id, post_id) VALUES
  ('ccccccc1-cccc-cccc-cccc-ccccccccccc1', 'Great post, Alice!', '22222222-2222-2222-2222-222222222222', 'aaaaaaa1-aaaa-aaaa-aaaa-aaaaaaaaaaa1'),
  ('ccccccc2-cccc-cccc-cccc-ccccccccccc2', 'Thanks for sharing!', '33333333-3333-3333-3333-333333333333', 'aaaaaaa1-aaaa-aaaa-aaaa-aaaaaaaaaaa1'),
  ('ccccccc3-cccc-cccc-cccc-ccccccccccc3', 'Interesting insights on tech.', '11111111-1111-1111-1111-111111111111', 'aaaaaaa3-aaaa-aaaa-aaaa-aaaaaaaaaaa3');
