create table if not exists users
(
  id serial primary key,
  email varchar(255) unique not null,
  password varchar(255) not null,
  name varchar(255) not null,
  birth_date date not null,
  gender varchar(10) not null check (gender IN ('MALE', 'FEMALE', 'OTHER')),
  location varchar(255) not null,
  profile_picture_url varchar(255),
  created_at timestamp with time zone not null default current_timestamp,
  updated_at timestamp with time zone not null default current_timestamp
);

create table if not exists profiles
(
  id serial primary key,
  user_id integer not null references users(id),
  bio text,
  interests text,
  verified boolean not null default false,
  created_at timestamp with time zone not null default current_timestamp,
  updated_at timestamp with time zone not null default current_timestamp
);

create table if not exists activities
(
  id serial primary key,
  user_id integer not null references users(id),
  profile_id integer not null references profiles(id),
  action varchar(10) not null check (action IN ('LIKE', 'PASS')),
  created_at timestamp with time zone not null default current_timestamp
);

create table if not exists activity_counters
(
  id serial primary key,
  user_id integer not null references users(id),
  count integer not null default 0,
  date date not null,
  created_at timestamp with time zone not null default current_timestamp,
  updated_at timestamp with time zone not null default current_timestamp
);

create table if not exists subscription_packages
(
  id serial primary key,
  name varchar(255) unique not null,
  lifetime_days integer not null default 1
);

create table if not exists subscriptions
(
  id serial primary key,
  subscription_package_id integer not null references subscription_packages(id),
  user_id integer not null references users(id),
  start_date date not null,
  end_date date not null,
  active boolean not null default false
);
