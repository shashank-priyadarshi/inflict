CREATE TABLE IF NOT EXISTS Members (
    id UUID NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type member_type NOT NULL,
    net_worth_id UUID,
    deleted BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_members_net_worth_id FOREIGN KEY (net_worth_id) REFERENCES Worths(id)
);

CREATE INDEX idx_members_net_worth_id ON Members(net_worth_id);

CREATE TABLE IF NOT EXISTS Amounts (
    id UUID NOT NULL PRIMARY KEY,
    type amount_type NOT NULL,
    name VARCHAR(255) NOT NULL,
    sender VARCHAR(100),
    receiver VARCHAR(100),
    value NUMERIC(20, 2) NOT NULL,
    currency VARCHAR(3) NOT NULL CHECK (char_length(currency) = 3),
    deleted BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS Worths (
    id UUID PRIMARY KEY,
    deleted BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS Wealths (
    id UUID NOT NULL PRIMARY KEY,
    worth_id UUID,
    type wealth_type NOT NULL,
    name VARCHAR(255) NOT NULL,
    value_id UUID NOT NULL,
    deleted BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_wealths_worth_id FOREIGN KEY (worth_id) REFERENCES Worths(id),
    CONSTRAINT fk_wealths_value_id FOREIGN KEY (value_id) REFERENCES Amounts(id)
);

CREATE TABLE IF NOT EXISTS Returns (
    id UUID NOT NULL PRIMARY KEY,
    wealth_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    rate_type rate_type NOT NULL,
    rate_value NUMERIC(20, 4) NOT NULL,
    duration INTERVAL,
    maturity_corpus_id UUID,
    deleted BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_returns_wealths_id FOREIGN KEY (wealth_id) REFERENCES Wealths(id),
    CONSTRAINT fk_returns_maturity_corpus_id FOREIGN KEY (maturity_corpus_id) REFERENCES Amounts(id)
);

CREATE TABLE IF NOT EXISTS Maintainances (
    id UUID NOT NULL PRIMARY KEY,
    wealth_id UUID NOT NULL,
    type VARCHAR(50) NOT NULL,
    name VARCHAR(255) NOT NULL,
    cost_id UUID NOT NULL,
    deleted BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_maintainances_wealth_id FOREIGN KEY (wealth_id) REFERENCES Wealths(id),
    CONSTRAINT fk_maintainances_cost_id FOREIGN KEY (cost_id) REFERENCES Amounts(id)
);

CREATE TABLE IF NOT EXISTS Accomodations (
    id UUID NOT NULL PRIMARY KEY,
    member_id UUID NOT NULL,
    type accomodation_type NOT NULL,
    address VARCHAR(255),
    cost_id UUID NOT NULL,
    deleted BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_accomodations_member_id FOREIGN KEY (member_id) REFERENCES Members(id),
    CONSTRAINT fk_accomodations_cost_id FOREIGN KEY (cost_id) REFERENCES Amounts(id)
);

CREATE INDEX idx_wealths_worth_id ON Wealths(worth_id);
CREATE INDEX idx_wealths_type ON Wealths(type);
CREATE INDEX idx_returns_wealth_id ON Returns(wealth_id);
CREATE INDEX idx_maintainances_wealth_id ON Maintainances(wealth_id);
CREATE INDEX idx_accomodations_member_id ON Accomodations(member_id);
