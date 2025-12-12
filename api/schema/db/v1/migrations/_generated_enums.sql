CREATE TYPE accomodation_type AS ENUM (
    'Owned',
    'Leased',
    'Rented',
    'Shared'
);

CREATE TYPE amount_type AS ENUM (
    'Credit',
    'Debit'
);

CREATE TYPE entity_type AS ENUM (
    'ENTITY_AMOUNT',
    'ENTITY_MAINTAINANCE',
    'ENTITY_RETURN',
    'ENTITY_WEALTH',
    'ENTITY_WORTH'
);

CREATE TYPE family_type AS ENUM (
    'Nuclear',
    'Joint'
);

CREATE TYPE member_type AS ENUM (
    'Earner',
    'Dependent'
);

CREATE TYPE rate_type AS ENUM (
    'BASIS_POINTS',
    'PERCENTAGE'
);

CREATE TYPE wealth_type AS ENUM (
    'Earning',
    'Expense',
    'Liability',
    'Saving',
    'Investment',
    'Insurance'
);
