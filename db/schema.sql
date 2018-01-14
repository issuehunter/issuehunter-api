DROP TABLE IF EXISTS campaigns;

CREATE TABLE IF NOT EXISTS campaigns(
  uid VARCHAR(255) NOT NULL PRIMARY KEY,
  rewarded BOOLEAN,
  total NUMERIC(100), -- uint256
  created_by VARCHAR(255),
  pre_reward_period_expires_at TIMESTAMP,
  reward_period_expires_at TIMESTAMP,
  resolved_by VARCHAR(255),
  patch_verifier VARCHAR(255),
  tip_per_mille INT,
  tips_amount NUMERIC(100),
  repo_url VARCHAR(255),
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  created_at_block_id VARCHAR(255) NOT NULL,
  updated_at_block_id VARCHAR(255) NOT NULL
);

CREATE INDEX campaigns_total_idx ON campaigns (total);
CREATE INDEX campaigns_created_at_idx ON campaigns (created_at);
CREATE INDEX campaigns_updated_at_idx ON campaigns (updated_at);

DROP TABLE IF EXISTS funds;

CREATE TABLE IF NOT EXISTS funds(
  funder_address VARCHAR(255),
  amount NUMERIC(100), -- uint256
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  created_at_block_id VARCHAR(255) NOT NULL,
  updated_at_block_id VARCHAR(255) NOT NULL
);

CREATE INDEX funder_address_idx ON funds (funder_address);
CREATE INDEX funds_created_at_idx ON funds (created_at);
CREATE INDEX funds_updated_at_idx ON funds (updated_at);

DROP TABLE IF EXISTS patches;

CREATE TABLE IF NOT EXISTS patches(
  author_address VARCHAR(255),
  ref VARCHAR(255),
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  created_at_block_id VARCHAR(255) NOT NULL,
  updated_at_block_id VARCHAR(255) NOT NULL
);

CREATE INDEX author_address_idx ON funds (author_address);
CREATE INDEX patches_created_at_idx ON patches (created_at);
CREATE INDEX patches_updated_at_idx ON patches (updated_at);
