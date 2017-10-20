--
-- PostgreSQL database dump
--

-- Dumped from database version 10.0
-- Dumped by pg_dump version 10.0

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner:
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner:
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: app_platform; Type: TABLE; Schema: public; Owner: hermes
--

CREATE TABLE app_platform (
    app_id integer NOT NULL,
    platform_id integer NOT NULL
);


ALTER TABLE app_platform OWNER TO hermes;

--
-- Name: app_user; Type: TABLE; Schema: public; Owner: hermes
--

CREATE TABLE app_user (
    app_id integer NOT NULL,
    user_id integer NOT NULL
);


ALTER TABLE app_user OWNER TO hermes;

--
-- Name: app_user_app; Type: TABLE; Schema: public; Owner: hermes
--

CREATE TABLE app_user_app (
    app_user_id integer NOT NULL,
    app_id integer NOT NULL
);


ALTER TABLE app_user_app OWNER TO hermes;

--
-- Name: app_user_device; Type: TABLE; Schema: public; Owner: hermes
--

CREATE TABLE app_user_device (
    app_user_id integer NOT NULL,
    device_id integer NOT NULL
);


ALTER TABLE app_user_device OWNER TO hermes;

--
-- Name: app_user_platform; Type: TABLE; Schema: public; Owner: hermes
--

CREATE TABLE app_user_platform (
    app_user_id integer NOT NULL,
    platform_id integer NOT NULL
);


ALTER TABLE app_user_platform OWNER TO hermes;

--
-- Name: apps; Type: TABLE; Schema: public; Owner: hermes
--

CREATE TABLE apps (
    id integer NOT NULL,
    name character varying(50) NOT NULL,
    type character(1) NOT NULL,
    key character(32) NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    updated_by integer,
    deleted_at timestamp(0) without time zone
);


ALTER TABLE apps OWNER TO hermes;

--
-- Name: apps_id_seq; Type: SEQUENCE; Schema: public; Owner: hermes
--

CREATE SEQUENCE apps_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE apps_id_seq OWNER TO hermes;

--
-- Name: apps_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hermes
--

ALTER SEQUENCE apps_id_seq OWNED BY apps.id;


--
-- Name: appusers; Type: TABLE; Schema: public; Owner: hermes
--

CREATE TABLE appusers (
    id bigint NOT NULL,
    name character varying(70) NOT NULL,
    email character varying(100),
    miba_id uuid,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    deleted_at timestamp(0) without time zone
);


ALTER TABLE appusers OWNER TO hermes;

--
-- Name: appusers_id_seq; Type: SEQUENCE; Schema: public; Owner: hermes
--

CREATE SEQUENCE appusers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE appusers_id_seq OWNER TO hermes;

--
-- Name: appusers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hermes
--

ALTER SEQUENCE appusers_id_seq OWNED BY appusers.id;


--
-- Name: brands; Type: TABLE; Schema: public; Owner: hermes
--

CREATE TABLE brands (
    id integer NOT NULL,
    name character varying(30) NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    deleted_at timestamp(0) without time zone
);


ALTER TABLE brands OWNER TO hermes;

--
-- Name: brands_id_seq; Type: SEQUENCE; Schema: public; Owner: hermes
--

CREATE SEQUENCE brands_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE brands_id_seq OWNER TO hermes;

--
-- Name: brands_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hermes
--

ALTER SEQUENCE brands_id_seq OWNED BY brands.id;


--
-- Name: browsers; Type: TABLE; Schema: public; Owner: hermes
--

CREATE TABLE browsers (
    id integer NOT NULL,
    name character varying(15) NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    deleted_at timestamp(0) without time zone
);


ALTER TABLE browsers OWNER TO hermes;

--
-- Name: browsers_id_seq; Type: SEQUENCE; Schema: public; Owner: hermes
--

CREATE SEQUENCE browsers_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE browsers_id_seq OWNER TO hermes;

--
-- Name: browsers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hermes
--

ALTER SEQUENCE browsers_id_seq OWNED BY browsers.id;


--
-- Name: categories; Type: TABLE; Schema: public; Owner: hermes
--

CREATE TABLE categories (
    id integer NOT NULL,
    parent_id integer,
    "order" integer DEFAULT 1 NOT NULL,
    name character varying(255) NOT NULL,
    slug character varying(255) NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


ALTER TABLE categories OWNER TO hermes;

--
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: hermes
--

CREATE SEQUENCE categories_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE categories_id_seq OWNER TO hermes;

--
-- Name: categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hermes
--

ALTER SEQUENCE categories_id_seq OWNED BY categories.id;


--
-- Name: data_rows; Type: TABLE; Schema: public; Owner: hermes
--

CREATE TABLE data_rows (
    id integer NOT NULL,
    data_type_id integer NOT NULL,
    field character varying(255) NOT NULL,
    type character varying(255) NOT NULL,
    display_name character varying(255) NOT NULL,
    required boolean DEFAULT false NOT NULL,
    browse boolean DEFAULT true NOT NULL,
    read boolean DEFAULT true NOT NULL,
    edit boolean DEFAULT true NOT NULL,
    add boolean DEFAULT true NOT NULL,
    delete boolean DEFAULT true NOT NULL,
    details text,
    "order" integer DEFAULT 1 NOT NULL
);


ALTER TABLE data_rows OWNER TO hermes;

--
-- Name: data_rows_id_seq; Type: SEQUENCE; Schema: public; Owner: hermes
--

CREATE SEQUENCE data_rows_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE data_rows_id_seq OWNER TO hermes;

--
-- Name: data_rows_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hermes
--

ALTER SEQUENCE data_rows_id_seq OWNED BY data_rows.id;


--
-- Name: data_types; Type: TABLE; Schema: public; Owner: hermes
--

CREATE TABLE data_types (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    slug character varying(255) NOT NULL,
    display_name_singular character varying(255) NOT NULL,
    display_name_plural character varying(255) NOT NULL,
    icon character varying(255),
    model_name character varying(255),
    description character varying(255),
    generate_permissions boolean DEFAULT false NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    server_side smallint DEFAULT '0'::smallint NOT NULL,
    controller character varying(255),
    policy_name character varying(255)
);


ALTER TABLE data_types OWNER TO hermes;

--
-- Name: data_types_id_seq; Type: SEQUENCE; Schema: public; Owner: hermes
--

CREATE SEQUENCE data_types_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE data_types_id_seq OWNER TO hermes;

--
-- Name: data_types_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hermes
--

ALTER SEQUENCE data_types_id_seq OWNED BY data_types.id;


--
-- Name: devices; Type: TABLE; Schema: public; Owner: hermes
--

CREATE TABLE devices (
    id integer NOT NULL,
    name character varying(63) NOT NULL,
    screen_width integer NOT NULL,
    screen_height integer NOT NULL,
    ppi integer,
    brand_id integer,
    platform_id integer NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    deleted_at timestamp(0) without time zone
);


ALTER TABLE devices OWNER TO hermes;

--
-- Name: devices_id_seq; Type: SEQUENCE; Schema: public; Owner: hermes
--

CREATE SEQUENCE devices_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE devices_id_seq OWNER TO hermes;

--
-- Name: devices_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hermes
--

ALTER SEQUENCE devices_id_seq OWNED BY devices.id;


--
-- Name: failed_jobs; Type: TABLE; Schema: public; Owner: hermes
--

CREATE TABLE failed_jobs (
    id bigint NOT NULL,
    connection text NOT NULL,
    queue text NOT NULL,
    payload text NOT NULL,
    exception text NOT NULL,
    failed_at timestamp(0) without time zone DEFAULT CURRENT_TIMESTAMP(0) NOT NULL
);


ALTER TABLE failed_jobs OWNER TO hermes;

--
-- Name: failed_jobs_id_seq; Type: SEQUENCE; Schema: public; Owner: hermes
--

CREATE SEQUENCE failed_jobs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE failed_jobs_id_seq OWNER TO hermes;

--
-- Name: failed_jobs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hermes
--

ALTER SEQUENCE failed_jobs_id_seq OWNED BY failed_jobs.id;


--
-- Name: menu_items; Type: TABLE; Schema: public; Owner: hermes
--

CREATE TABLE menu_items (
    id integer NOT NULL,
    menu_id integer,
    title character varying(255) NOT NULL,
    url character varying(255) NOT NULL,
    target character varying(255) DEFAULT '_self'::character varying NOT NULL,
    icon_class character varying(255),
    color character varying(255),
    parent_id integer,
    "order" integer NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    route character varying(255),
    parameters text
);


ALTER TABLE menu_items OWNER TO hermes;

--
-- Name: menu_items_id_seq; Type: SEQUENCE; Schema: public; Owner: hermes
--

CREATE SEQUENCE menu_items_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE menu_items_id_seq OWNER TO hermes;

--
-- Name: menu_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hermes
--

ALTER SEQUENCE menu_items_id_seq OWNED BY menu_items.id;


--
-- Name: menus; Type: TABLE; Schema: public; Owner: hermes
--

CREATE TABLE menus (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


ALTER TABLE menus OWNER TO hermes;

--
-- Name: menus_id_seq; Type: SEQUENCE; Schema: public; Owner: hermes
--

CREATE SEQUENCE menus_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE menus_id_seq OWNER TO hermes;

--
-- Name: menus_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hermes
--

ALTER SEQUENCE menus_id_seq OWNED BY menus.id;


--
-- Name: messages; Type: TABLE; Schema: public; Owner: hermes
--

CREATE TABLE messages (
    id bigint NOT NULL,
    message text NOT NULL,
    direction character varying(255) NOT NULL,
    status smallint DEFAULT '0'::smallint NOT NULL,
    transport_id character varying(90),
    rating_id integer NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    CONSTRAINT messages_direction_check CHECK (((direction)::text = ANY ((ARRAY['in'::character varying, 'out'::character varying])::text[])))
);


ALTER TABLE messages OWNER TO hermes;

--
-- Name: messages_id_seq; Type: SEQUENCE; Schema: public; Owner: hermes
--

CREATE SEQUENCE messages_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE messages_id_seq OWNER TO hermes;

--
-- Name: messages_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hermes
--

ALTER SEQUENCE messages_id_seq OWNED BY messages.id;


--
-- Name: migrations; Type: TABLE; Schema: public; Owner: hermes
--

CREATE TABLE migrations (
    id integer NOT NULL,
    migration character varying(255) NOT NULL,
    batch integer NOT NULL
);


ALTER TABLE migrations OWNER TO hermes;

--
-- Name: migrations_id_seq; Type: SEQUENCE; Schema: public; Owner: hermes
--

CREATE SEQUENCE migrations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE migrations_id_seq OWNER TO hermes;

--
-- Name: migrations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hermes
--

ALTER SEQUENCE migrations_id_seq OWNED BY migrations.id;


--
-- Name: pages; Type: TABLE; Schema: public; Owner: hermes
--

CREATE TABLE pages (
    id integer NOT NULL,
    author_id integer NOT NULL,
    title character varying(255) NOT NULL,
    excerpt text,
    body text,
    image character varying(255),
    slug character varying(255) NOT NULL,
    meta_description text,
    meta_keywords text,
    status character varying(255) DEFAULT 'INACTIVE'::character varying NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    CONSTRAINT pages_status_check CHECK (((status)::text = ANY ((ARRAY['ACTIVE'::character varying, 'INACTIVE'::character varying])::text[])))
);


ALTER TABLE pages OWNER TO hermes;

--
-- Name: pages_id_seq; Type: SEQUENCE; Schema: public; Owner: hermes
--

CREATE SEQUENCE pages_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE pages_id_seq OWNER TO hermes;

--
-- Name: pages_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hermes
--

ALTER SEQUENCE pages_id_seq OWNED BY pages.id;


--
-- Name: password_resets; Type: TABLE; Schema: public; Owner: hermes
--

CREATE TABLE password_resets (
    email character varying(255) NOT NULL,
    token character varying(255) NOT NULL,
    created_at timestamp(0) without time zone
);


ALTER TABLE password_resets OWNER TO hermes;

--
-- Name: permission_groups; Type: TABLE; Schema: public; Owner: hermes
--

CREATE TABLE permission_groups (
    id integer NOT NULL,
    name character varying(255) NOT NULL
);


ALTER TABLE permission_groups OWNER TO hermes;

--
-- Name: permission_groups_id_seq; Type: SEQUENCE; Schema: public; Owner: hermes
--

CREATE SEQUENCE permission_groups_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE permission_groups_id_seq OWNER TO hermes;

--
-- Name: permission_groups_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hermes
--

ALTER SEQUENCE permission_groups_id_seq OWNED BY permission_groups.id;


--
-- Name: permission_role; Type: TABLE; Schema: public; Owner: hermes
--

CREATE TABLE permission_role (
    permission_id integer NOT NULL,
    role_id integer NOT NULL
);


ALTER TABLE permission_role OWNER TO hermes;

--
-- Name: permissions; Type: TABLE; Schema: public; Owner: hermes
--

CREATE TABLE permissions (
    id integer NOT NULL,
    key character varying(255) NOT NULL,
    table_name character varying(255),
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    permission_group_id integer
);


ALTER TABLE permissions OWNER TO hermes;

--
-- Name: permissions_id_seq; Type: SEQUENCE; Schema: public; Owner: hermes
--

CREATE SEQUENCE permissions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE permissions_id_seq OWNER TO hermes;

--
-- Name: permissions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hermes
--

ALTER SEQUENCE permissions_id_seq OWNED BY permissions.id;


--
-- Name: platforms; Type: TABLE; Schema: public; Owner: hermes
--

CREATE TABLE platforms (
    id integer NOT NULL,
    name character varying(15) NOT NULL,
    key character(32) NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    deleted_at timestamp(0) without time zone
);


ALTER TABLE platforms OWNER TO hermes;

--
-- Name: platforms_id_seq; Type: SEQUENCE; Schema: public; Owner: hermes
--

CREATE SEQUENCE platforms_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE platforms_id_seq OWNER TO hermes;

--
-- Name: platforms_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hermes
--

ALTER SEQUENCE platforms_id_seq OWNED BY platforms.id;


--
-- Name: posts; Type: TABLE; Schema: public; Owner: hermes
--

CREATE TABLE posts (
    id integer NOT NULL,
    author_id integer NOT NULL,
    category_id integer,
    title character varying(255) NOT NULL,
    seo_title character varying(255),
    excerpt text,
    body text NOT NULL,
    image character varying(255),
    slug character varying(255) NOT NULL,
    meta_description text,
    meta_keywords text,
    status character varying(255) DEFAULT 'DRAFT'::character varying NOT NULL,
    featured boolean DEFAULT false NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    CONSTRAINT posts_status_check CHECK (((status)::text = ANY ((ARRAY['PUBLISHED'::character varying, 'DRAFT'::character varying, 'PENDING'::character varying])::text[])))
);


ALTER TABLE posts OWNER TO hermes;

--
-- Name: posts_id_seq; Type: SEQUENCE; Schema: public; Owner: hermes
--

CREATE SEQUENCE posts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE posts_id_seq OWNER TO hermes;

--
-- Name: posts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hermes
--

ALTER SEQUENCE posts_id_seq OWNED BY posts.id;


--
-- Name: ranges; Type: TABLE; Schema: public; Owner: hermes
--

CREATE TABLE ranges (
    id integer NOT NULL,
    name character varying(11) NOT NULL,
    "from" integer NOT NULL,
    "to" integer NOT NULL,
    key character(32) NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    deleted_at timestamp(0) without time zone
);


ALTER TABLE ranges OWNER TO hermes;

--
-- Name: ranges_id_seq; Type: SEQUENCE; Schema: public; Owner: hermes
--

CREATE SEQUENCE ranges_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE ranges_id_seq OWNER TO hermes;

--
-- Name: ranges_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hermes
--

ALTER SEQUENCE ranges_id_seq OWNED BY ranges.id;


--
-- Name: ratings; Type: TABLE; Schema: public; Owner: hermes
--

CREATE TABLE ratings (
    id bigint NOT NULL,
    rating smallint NOT NULL,
    description character varying(30),
    app_version character varying(15),
    browser_version character varying(15),
    platform_version character varying(15) NOT NULL,
    has_message boolean DEFAULT false NOT NULL,
    app_id integer NOT NULL,
    range_id integer NOT NULL,
    platform_id integer NOT NULL,
    device_id integer NOT NULL,
    appuser_id integer,
    browser_id integer,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    deleted_at timestamp(0) without time zone
);


ALTER TABLE ratings OWNER TO hermes;

--
-- Name: ratings_id_seq; Type: SEQUENCE; Schema: public; Owner: hermes
--

CREATE SEQUENCE ratings_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE ratings_id_seq OWNER TO hermes;

--
-- Name: ratings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hermes
--

ALTER SEQUENCE ratings_id_seq OWNED BY ratings.id;


--
-- Name: roles; Type: TABLE; Schema: public; Owner: hermes
--

CREATE TABLE roles (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    display_name character varying(255) NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


ALTER TABLE roles OWNER TO hermes;

--
-- Name: roles_id_seq; Type: SEQUENCE; Schema: public; Owner: hermes
--

CREATE SEQUENCE roles_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE roles_id_seq OWNER TO hermes;

--
-- Name: roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hermes
--

ALTER SEQUENCE roles_id_seq OWNED BY roles.id;


--
-- Name: settings; Type: TABLE; Schema: public; Owner: hermes
--

CREATE TABLE settings (
    id integer NOT NULL,
    key character varying(255) NOT NULL,
    display_name character varying(255) NOT NULL,
    value text NOT NULL,
    details text,
    type character varying(255) NOT NULL,
    "order" integer DEFAULT 1 NOT NULL,
    "group" character varying(255)
);


ALTER TABLE settings OWNER TO hermes;

--
-- Name: settings_id_seq; Type: SEQUENCE; Schema: public; Owner: hermes
--

CREATE SEQUENCE settings_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE settings_id_seq OWNER TO hermes;

--
-- Name: settings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hermes
--

ALTER SEQUENCE settings_id_seq OWNED BY settings.id;


--
-- Name: translations; Type: TABLE; Schema: public; Owner: hermes
--

CREATE TABLE translations (
    id integer NOT NULL,
    table_name character varying(255) NOT NULL,
    column_name character varying(255) NOT NULL,
    foreign_key integer NOT NULL,
    locale character varying(255) NOT NULL,
    value text NOT NULL,
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


ALTER TABLE translations OWNER TO hermes;

--
-- Name: translations_id_seq; Type: SEQUENCE; Schema: public; Owner: hermes
--

CREATE SEQUENCE translations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE translations_id_seq OWNER TO hermes;

--
-- Name: translations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hermes
--

ALTER SEQUENCE translations_id_seq OWNED BY translations.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: hermes
--

CREATE TABLE users (
    id integer NOT NULL,
    name character varying(70) NOT NULL,
    email character varying(100) NOT NULL,
    remember_token character varying(100),
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone,
    updated_by integer,
    avatar character varying(255),
    role_id integer
);


ALTER TABLE users OWNER TO hermes;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: hermes
--

CREATE SEQUENCE users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE users_id_seq OWNER TO hermes;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: hermes
--

ALTER SEQUENCE users_id_seq OWNED BY users.id;


--
-- Name: apps id; Type: DEFAULT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY apps ALTER COLUMN id SET DEFAULT nextval('apps_id_seq'::regclass);


--
-- Name: appusers id; Type: DEFAULT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY appusers ALTER COLUMN id SET DEFAULT nextval('appusers_id_seq'::regclass);


--
-- Name: brands id; Type: DEFAULT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY brands ALTER COLUMN id SET DEFAULT nextval('brands_id_seq'::regclass);


--
-- Name: browsers id; Type: DEFAULT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY browsers ALTER COLUMN id SET DEFAULT nextval('browsers_id_seq'::regclass);


--
-- Name: categories id; Type: DEFAULT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY categories ALTER COLUMN id SET DEFAULT nextval('categories_id_seq'::regclass);


--
-- Name: data_rows id; Type: DEFAULT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY data_rows ALTER COLUMN id SET DEFAULT nextval('data_rows_id_seq'::regclass);


--
-- Name: data_types id; Type: DEFAULT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY data_types ALTER COLUMN id SET DEFAULT nextval('data_types_id_seq'::regclass);


--
-- Name: devices id; Type: DEFAULT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY devices ALTER COLUMN id SET DEFAULT nextval('devices_id_seq'::regclass);


--
-- Name: failed_jobs id; Type: DEFAULT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY failed_jobs ALTER COLUMN id SET DEFAULT nextval('failed_jobs_id_seq'::regclass);


--
-- Name: menu_items id; Type: DEFAULT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY menu_items ALTER COLUMN id SET DEFAULT nextval('menu_items_id_seq'::regclass);


--
-- Name: menus id; Type: DEFAULT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY menus ALTER COLUMN id SET DEFAULT nextval('menus_id_seq'::regclass);


--
-- Name: messages id; Type: DEFAULT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY messages ALTER COLUMN id SET DEFAULT nextval('messages_id_seq'::regclass);


--
-- Name: migrations id; Type: DEFAULT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY migrations ALTER COLUMN id SET DEFAULT nextval('migrations_id_seq'::regclass);


--
-- Name: pages id; Type: DEFAULT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY pages ALTER COLUMN id SET DEFAULT nextval('pages_id_seq'::regclass);


--
-- Name: permission_groups id; Type: DEFAULT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY permission_groups ALTER COLUMN id SET DEFAULT nextval('permission_groups_id_seq'::regclass);


--
-- Name: permissions id; Type: DEFAULT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY permissions ALTER COLUMN id SET DEFAULT nextval('permissions_id_seq'::regclass);


--
-- Name: platforms id; Type: DEFAULT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY platforms ALTER COLUMN id SET DEFAULT nextval('platforms_id_seq'::regclass);


--
-- Name: posts id; Type: DEFAULT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY posts ALTER COLUMN id SET DEFAULT nextval('posts_id_seq'::regclass);


--
-- Name: ranges id; Type: DEFAULT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY ranges ALTER COLUMN id SET DEFAULT nextval('ranges_id_seq'::regclass);


--
-- Name: ratings id; Type: DEFAULT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY ratings ALTER COLUMN id SET DEFAULT nextval('ratings_id_seq'::regclass);


--
-- Name: roles id; Type: DEFAULT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY roles ALTER COLUMN id SET DEFAULT nextval('roles_id_seq'::regclass);


--
-- Name: settings id; Type: DEFAULT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY settings ALTER COLUMN id SET DEFAULT nextval('settings_id_seq'::regclass);


--
-- Name: translations id; Type: DEFAULT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY translations ALTER COLUMN id SET DEFAULT nextval('translations_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY users ALTER COLUMN id SET DEFAULT nextval('users_id_seq'::regclass);


--
-- Data for Name: app_platform; Type: TABLE DATA; Schema: public; Owner: hermes
--

COPY app_platform (app_id, platform_id) FROM stdin;
1	1
2	1
3	1
1	2
2	2
3	2
\.


--
-- Data for Name: app_user; Type: TABLE DATA; Schema: public; Owner: hermes
--

COPY app_user (app_id, user_id) FROM stdin;
1	2
2	2
3	2
1	3
2	3
3	3
1	4
2	4
3	4
\.


--
-- Data for Name: app_user_app; Type: TABLE DATA; Schema: public; Owner: hermes
--

COPY app_user_app (app_user_id, app_id) FROM stdin;
\.


--
-- Data for Name: app_user_device; Type: TABLE DATA; Schema: public; Owner: hermes
--

COPY app_user_device (app_user_id, device_id) FROM stdin;
1	1
2	1
3	1
4	1
5	1
6	1
1	2
2	2
3	2
4	2
5	2
6	2
1	3
2	3
3	3
4	3
5	3
6	3
1	4
2	4
3	4
4	4
5	4
6	4
1	5
2	5
3	5
4	5
5	5
6	5
\.


--
-- Data for Name: app_user_platform; Type: TABLE DATA; Schema: public; Owner: hermes
--

COPY app_user_platform (app_user_id, platform_id) FROM stdin;
1	1
2	1
4	1
3	2
5	2
6	2
\.


--
-- Data for Name: apps; Type: TABLE DATA; Schema: public; Owner: hermes
--

COPY apps (id, name, type, key, created_at, updated_at, updated_by, deleted_at) FROM stdin;
1	Denuncia Vial	M	e10adc3949ba59abbe56e057f20f883e	2017-10-11 20:25:59	\N	\N	\N
2	Mi BA	M	c33367701511b4f6020ec61ded352059	2017-10-11 20:25:59	\N	\N	\N
3	Más Simple	M	0d06fd8cb26eb57f1a690f493663cc55	2017-10-11 20:25:59	\N	\N	\N
\.


--
-- Data for Name: appusers; Type: TABLE DATA; Schema: public; Owner: hermes
--

COPY appusers (id, name, email, miba_id, created_at, updated_at, deleted_at) FROM stdin;
1	Mariano Gómez	mariano@gomez.com	412319c6-536d-4942-b5bc-0edbd696b9d7	2017-10-11 20:25:59	\N	\N
2	Esteban Sosa	esteban@sosa.com	88bf0e47-666b-468d-91b0-b0eb5955298a	2017-10-11 20:25:59	\N	\N
3	Germán Álvarez	german@alvarez.com	dfe21bc0-47ef-40ef-9772-f91c57948510	2017-10-11 20:25:59	\N	\N
4	Mariela Domínguez	mariela@dominguez.com	1cb2d373-f265-4e76-a853-a56a1bf30a0e	2017-10-11 20:25:59	\N	\N
5	Juliana Pérez	juliana@perez.com	6afb1806-9e44-4622-894e-20538e3a33cf	2017-10-11 20:25:59	\N	\N
6	Valentina Echeverría	valentina@echeverria.com	b8051abe-1a03-4023-9a54-a2ea79c1f8dc	2017-10-11 20:25:59	\N	\N
\.


--
-- Data for Name: brands; Type: TABLE DATA; Schema: public; Owner: hermes
--

COPY brands (id, name, created_at, updated_at, deleted_at) FROM stdin;
1	Apple	2017-10-11 20:25:59	\N	\N
2	Samsung	2017-10-11 20:25:59	\N	\N
3	Google	2017-10-11 20:25:59	\N	\N
\.


--
-- Data for Name: browsers; Type: TABLE DATA; Schema: public; Owner: hermes
--

COPY browsers (id, name, created_at, updated_at, deleted_at) FROM stdin;
1	Chrome	2017-10-11 20:25:59	\N	\N
2	Firefox	2017-10-11 20:25:59	\N	\N
3	Opera	2017-10-11 20:25:59	\N	\N
\.


--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: hermes
--

COPY categories (id, parent_id, "order", name, slug, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: data_rows; Type: TABLE DATA; Schema: public; Owner: hermes
--

COPY data_rows (id, data_type_id, field, type, display_name, required, browse, read, edit, add, delete, details, "order") FROM stdin;
1	11	id	number	ID	t	f	f	f	f	f		1
2	11	name	text	Nombre	t	t	t	t	t	t		2
3	11	created_at	timestamp	Creación	f	f	f	f	f	f		3
4	11	updated_at	timestamp	Última Modificación	f	f	f	f	f	f		4
5	12	id	number	ID	t	f	f	f	f	f		1
6	12	name	text	Nombre	t	t	t	t	t	t		2
7	12	created_at	timestamp	Creación	f	f	f	f	f	f		3
8	12	updated_at	timestamp	Última Modificación	f	f	f	f	f	f		4
9	12	display_name	text	Display Name	t	t	t	t	t	t		5
10	2	id	number	ID	t	f	f	f	f	f		1
11	2	name	text	Nombre	t	t	t	t	t	t	{"validation":{"rules":["required","string","min:3","max:50"],"messages":{"required":"Falta el campo :attribute.","string":"El campo :attribute debe tener texto.","max":"El campo :attribute puede tener hasta :max carácteres.","min":"El campo :attribute debe tener al menos :min carácteres."}}}	2
12	2	type	radio_btn	Tipo	t	t	t	t	t	t	{"default":"M","options":{"Móvil":"Móvil","Web":"Web"},"validation":{"rules":["required","alpha","size:1"],"messages":{"required":"Falta el campo :attribute.","alpha":"El campo :attribute sólo puede constar de una letra.","size":"El campo :attribute sólo puede constar de una letra."}}}	3
13	2	platforms	select_multiple	Plataformas	t	f	t	t	t	t	{"relationship":{"key":"id","label":"name","page_slug":"admin/platforms"}}	4
14	2	users	select_multiple	Personal	t	f	t	t	t	t	{"relationship":{"key":"id","label":"name","page_slug":"admin/users"}}	5
15	2	key	text	Key	f	f	t	f	f	f		6
16	2	updated_at	timestamp	Última Modificación	f	f	t	f	f	f		7
17	2	updated_by	text	Modificado Por	f	f	t	f	f	f		8
18	2	created_at	timestamp	Creación	f	f	t	f	f	f		9
19	2	deleted_at	timestamp	Borrado	f	f	f	f	f	f		10
20	3	id	number	ID	t	f	f	f	f	f		1
21	3	name	text	Nombre	t	t	t	f	f	f	{"validation":{"rules":["required","string","min:3","max:70"],"messages":{"required":"Falta el campo :attribute.","string":"El campo :attribute debe tener texto.","max":"El campo :attribute puede tener hasta :max carácteres.","min":"El campo :attribute debe tener al menos :min carácteres."}}}	2
22	3	email	text	Email	t	t	t	f	f	f	{"validation":{"rules":["required","string","email","min:3","max:100"],"messages":{"required":"Falta el campo :attribute.","string":"El campo :attribute debe tener texto.","email":"El campo :attribute debe ser un email válido.","max":"El campo :attribute puede tener hasta :max carácteres.","min":"El campo :attribute debe tener al menos :min carácteres."}}}	3
23	3	ratings	select_dropdown	Calificaciones	t	f	t	f	f	f	{"relationship":{"key":"id","label":"rating","page_slug":"admin/ratings"}}	4
24	3	apps	select_multiple	Aplicaciones	t	f	t	f	f	f	{"relationship":{"key":"id","label":"name","page_slug":"admin/apps"}}	5
25	3	platforms	select_multiple	Plataformas	t	f	t	f	f	f	{"relationship":{"key":"id","label":"name","page_slug":"admin/platforms"}}	6
26	3	devices	select_multiple	Dispositivos	t	f	t	f	f	f	{"relationship":{"key":"id","label":"name","page_slug":"admin/devices"}}	7
27	3	miba_id	text	ID MiBA	t	t	t	f	f	f	{"validation":{"rules":["required","string"],"messages":{"required":"Falta el campo :attribute.","string":"El campo :attribute debe tener texto."}}}	8
28	3	updated_at	timestamp	Última Modificación	f	f	t	f	f	f		9
29	3	created_at	timestamp	Creación	f	t	t	f	f	f		10
30	3	deleted_at	timestamp	Borrado	f	f	f	f	f	f		11
31	4	id	number	ID	t	f	f	f	f	f		1
32	4	name	text	Nombre	t	t	t	f	f	f	{"validation":{"rules":["required","string","min:1","max:30"],"messages":{"required":"Falta el campo :attribute.","string":"El campo :attribute debe tener texto.","max":"El campo :attribute puede tener hasta :max carácteres.","min":"El campo :attribute debe tener al menos :min carácteres."}}}	2
33	4	updated_at	timestamp	Última Modificación	f	f	t	f	f	f		3
34	4	created_at	timestamp	Creación	f	t	t	f	f	f		4
35	4	deleted_at	timestamp	Borrado	f	f	f	f	f	f		5
36	5	id	number	ID	t	f	f	f	f	f		1
37	5	name	text	Nombre	t	t	t	f	f	f	{"validation":{"rules":["required","string","min:1","max:15"],"messages":{"required":"Falta el campo :attribute.","string":"El campo :attribute debe tener texto.","max":"El campo :attribute puede tener hasta :max carácteres.","min":"El campo :attribute debe tener al menos :min carácteres."}}}	2
38	5	updated_at	timestamp	Última Modificación	f	f	t	f	f	f		3
39	5	created_at	timestamp	Creación	f	t	t	f	f	f		4
40	5	deleted_at	timestamp	Borrado	f	f	f	f	f	f		5
41	8	id	number	ID	t	f	f	f	f	f		1
42	8	name	text	Nombre	t	t	t	t	t	t	{"validation":{"rules":["required","string","min:2","max:30"],"messages":{"required":"Falta el campo :attribute.","string":"El campo :attribute debe tener texto.","max":"El campo :attribute puede tener hasta :max carácteres.","min":"El campo :attribute debe tener al menos :min carácteres."}}}	2
43	8	key	text	Key	f	f	t	f	f	f		3
44	8	updated_at	timestamp	Última Modificación	f	f	t	f	f	f		4
45	8	created_at	timestamp	Creación	f	f	t	f	f	f		5
46	8	deleted_at	timestamp	Borrado	f	f	f	f	f	f		6
47	9	id	number	ID	t	f	f	f	f	f		1
48	9	name	text	Nombre	f	t	t	f	f	f		2
49	9	from	number	Desde	t	t	t	t	t	t	{"validation":{"rules":["required","integer","max:127","min:-127"],"messages":{"required":"Falta el campo :attribute.","integer":"El campo :attribute debe ser numérico.","max":"El campo :attribute puede ser hasta :max.","min":"El campo :attribute debe ser al menos :min."}}}	3
50	9	to	number	Hasta	t	t	t	t	t	t	{"validation":{"rules":["required","integer","max:127","min:-127"],"messages":{"required":"Falta el campo :attribute.","integer":"El campo :attribute debe ser numérico.","max":"El campo :attribute puede ser hasta :max.","min":"El campo :attribute debe ser al menos :min."}}}	4
51	9	key	text	Key	f	f	t	f	f	f		5
52	9	updated_at	timestamp	Última Modificación	f	f	t	f	f	f		6
53	9	created_at	timestamp	Creación	f	f	t	f	f	f		7
54	9	deleted_at	timestamp	Borrado	f	f	f	f	f	f		8
55	6	id	number	ID	t	f	f	f	f	f		1
56	6	brand_id	select_dropdown	Marca	t	t	t	f	f	f	{"validation":{"rules":["integer","nullable"],"messages":{"integer":"El campo :attribute debe ser un número entero."}},"relationship":{"key":"id","label":"name","page_slug":"admin/brands"}}	2
57	6	name	text	Nombre	t	t	t	f	f	f	{"validation":{"rules":["required","string","min:1","max:30"],"messages":{"required":"Falta el campo :attribute.","string":"El campo :attribute debe tener texto.","max":"El campo :attribute puede tener hasta :max carácteres.","min":"El campo :attribute debe tener al menos :min carácteres."}}}	3
58	6	screen_width	number	Ancho Pantalla	t	t	t	f	f	f	{"validation":{"rules":["required","integer","digits_between:3,5"],"messages":{"required":"Falta el campo :attribute.","integer":"El campo :attribute debe ser un número entero.","digits_between":"El campo :attribute debe estar entre :min y :max."}}}	4
59	6	screen_height	number	Altura Pantalla	t	t	t	f	f	f	{"validation":{"rules":["required","integer","digits_between:3,5"],"messages":{"required":"Falta el campo :attribute.","integer":"El campo :attribute debe ser un número entero.","digits_between":"El campo :attribute debe estar entre :min y :max."}}}	5
60	6	ppi	number	PPI	t	t	t	f	f	f	{"validation":{"rules":["required","integer","digits_between:3,4"],"messages":{"required":"Falta el campo :attribute.","integer":"El campo :attribute debe ser un número entero.","digits_between":"El campo :attribute debe estar entre :min y :max."}}}	6
61	6	platform_id	select_dropdown	Plataforma	t	t	t	f	f	f	{"validation":{"rules":["required","integer"],"messages":{"required":"Falta el campo :attribute.","integer":"El campo :attribute debe ser un número entero."}},"relationship":{"key":"id","label":"name","page_slug":"admin/platforms"}}	7
62	6	updated_at	timestamp	Última Modificación	f	f	t	f	f	f		8
63	6	created_at	timestamp	Creación	f	t	t	f	f	f		9
64	6	deleted_at	timestamp	Borrado	f	f	f	f	f	f		10
65	7	id	number	ID	t	f	f	f	f	f		1
66	7	message	text_area	Texto	t	t	t	f	t	t	{"validation":{"rules":["required","string","min:3","max:1000"],"messages":{"required":"Falta el campo :attribute.","string":"El campo :attribute debe tener texto.","max":"El campo :attribute puede tener hasta :max carácteres.","min":"El campo :attribute debe tener al menos :min carácteres."}}}	2
67	7	direction	text	Sentido	t	t	t	f	f	f	{"validation":{"rules":["required","string","in:in,out"],"messages":{"required":"Falta el campo :attribute.","string":"El campo :attribute debe tener texto.","in":"El valor del campo :attribute sólo puede ser 'in' u 'out'"}}}	3
68	7	rating_id	select_dropdown	Rating	t	t	t	f	t	t	{"validation":{"rules":["required","integer"],"messages":{"required":"Falta el campo :attribute.","integer":"El campo :attribute debe ser un número entero."}},"relationship":{"key":"id","label":"rating","page_slug":"admin/ratings"}}	4
69	7	updated_at	timestamp	Última Modificación	f	f	t	f	f	f		5
70	7	created_at	timestamp	Fecha	f	t	t	f	f	f		6
71	10	id	number	ID	t	f	f	f	f	f		1
72	10	rating	number	Calificación	t	t	t	f	f	f	{"validation":{"rules":["required","integer","min:-127","max:127"],"messages":{"required":"Falta el campo :attribute.","integer":"El campo :attribute debe ser un número entero.","max":"El campo :attribute puede ser hasta :max.","min":"El campo :attribute no debe ser menor a :min."}}}	2
73	10	range_id	select_dropdown	Rango	t	t	t	f	f	f	{"validation":{"rules":["required","integer"],"messages":{"required":"Falta el campo :attribute.","integer":"El campo :attribute debe ser un número entero."}},"relationship":{"key":"id","label":"name","page_slug":"admin/ranges"}}	3
74	10	description	text	Descripción	f	t	t	f	f	f	{"validation":{"rules":["string","min:1","max:30","nullable"],"messages":{"required":"Falta el campo :attribute.","string":"El campo :attribute debe tener texto.","max":"El campo :attribute puede tener hasta :max carácteres.","min":"El campo :attribute debe tener al menos :min carácteres."}}}	4
75	10	has_message	check	Mensaje	t	t	t	f	f	f	{"validation":{"rules":["required","boolean"],"messages":{"required":"Falta el campo :attribute.","boolean":"El campo :attribute debe ser verdadero o falso."}}}	5
76	10	app_id	select_dropdown	App	t	t	t	f	f	f	{"validation":{"rules":["required","integer"],"messages":{"required":"Falta el campo :attribute.","integer":"El campo :attribute debe ser un número entero."}},"relationship":{"key":"id","label":"name","page_slug":"admin/apps"}}	6
77	10	app_version	text	Versión	f	t	t	f	f	f	{"validation":{"rules":["string","min:1","max:15","nullable"],"messages":{"string":"El campo :attribute debe tener texto.","max":"El campo :attribute puede tener hasta :max carácteres.","min":"El campo :attribute debe tener al menos :min carácteres."}}}	7
78	10	platform_id	select_dropdown	Plataforma	t	t	t	f	f	f	{"validation":{"rules":["required","integer"],"messages":{"required":"Falta el campo :attribute.","integer":"El campo :attribute debe ser un número entero."}},"relationship":{"key":"id","label":"name"}}	8
79	10	platform_version	text	Versión	f	t	t	f	f	f	{"validation":{"rules":["string","min:1","max:15","nullable"],"messages":{"string":"El campo :attribute debe tener texto.","max":"El campo :attribute puede tener hasta :max carácteres.","min":"El campo :attribute debe tener al menos :min carácteres."}}}	9
80	10	browser_id	select_dropdown	Browser	f	t	t	f	f	f	{"validation":{"rules":["integer","nullable"],"messages":{"integer":"El campo :attribute debe ser un número entero."}},"relationship":{"key":"id","label":"name"}}	10
81	10	browser_version	text	Versión	f	t	t	f	f	f	{"validation":{"rules":["string","min:1","max:15","nullable"],"messages":{"string":"El campo :attribute debe tener texto.","max":"El campo :attribute puede tener hasta :max carácteres.","min":"El campo :attribute debe tener al menos :min carácteres."}}}	11
82	10	appuser_id	select_dropdown	Usuario	f	t	t	f	f	f	{"validation":{"rules":["integer","nullable"],"messages":{"integer":"El campo :attribute debe ser un número entero."}},"relationship":{"key":"id","label":"name","page_slug":"admin/appusers"}}	12
83	10	device_id	select_dropdown	Dispositivo	f	t	t	f	f	f	{"validation":{"rules":["integer","nullable"],"messages":{"integer":"El campo :attribute debe ser un número entero."}},"relationship":{"key":"id","label":"name","page_slug":"admin/devices"}}	13
84	10	updated_at	timestamp	Última Modificación	f	f	t	f	f	f		14
85	10	created_at	timestamp	Creación	f	t	t	f	f	f		15
86	10	deleted_at	timestamp	Borrado	f	f	f	f	f	f		16
87	1	id	number	ID	t	f	f	f	f	f		1
88	1	name	text	Nombre	t	t	t	t	t	t	{"validation":{"rules":["required","string","min:2","max:70"],"messages":{"required":"Falta el campo :attribute.","string":"El campo :attribute debe tener texto.","max":"El campo :attribute puede tener hasta :max carácteres.","min":"El campo :attribute debe tener al menos :min carácteres."}}}	2
89	1	email	text	Email	t	t	t	t	t	t	{"validation":{"rules":["required","string","email"],"messages":{"required":"Falta el campo :attribute.","string":"El campo :attribute debe tener texto.","email":"El campo :attribute debe ser un email válido."}}}	3
90	1	password	password	Password	t	f	f	t	t	f	{"validation":{"rules":["required","string","min:8","max:70"],"messages":{"required":"Falta el campo :attribute.","string":"El campo :attribute debe tener texto.","max":"El campo :attribute puede tener hasta :max carácteres.","min":"El campo :attribute debe tener al menos :min carácteres."}}}	4
91	1	role_id	select_dropdown	Rol	t	t	t	t	t	t	{"validation":{"rules":["required","integer"],"messages":{"required":"Falta el campo :attribute.","integer":"El campo :attribute debe ser un número entero."},"relationship":{"key":"id","label":"display_name","page_slug":"admin/roles"}}}	6
92	1	apps	select_multiple	Aplicaciones	f	f	t	t	t	t	{"relationship":{"key":"id","label":"name","page_slug":"admin/apps"}}	7
93	1	remember_token	text	Recordar Token	f	f	f	f	f	f		8
94	1	updated_at	timestamp	Última Modificación	f	f	f	f	f	f		9
95	1	updated_by	text	Modificado Por	f	f	t	f	f	f		10
96	1	created_at	timestamp	Creación	f	f	t	f	f	f		11
\.


--
-- Data for Name: data_types; Type: TABLE DATA; Schema: public; Owner: hermes
--

COPY data_types (id, name, slug, display_name_singular, display_name_plural, icon, model_name, description, generate_permissions, created_at, updated_at, server_side, controller, policy_name) FROM stdin;
1	users	users	User	Users	voyager-person	App\\User	Personal	f	2017-10-11 20:25:59	2017-10-11 20:25:59	0	Controller	\N
2	apps	apps	App	Apps	voyager-categories	App\\App	Aplicaciones	f	2017-10-11 20:25:59	2017-10-11 20:25:59	0	Controller	\N
3	appusers	appusers	App User	App Users	voyager-people	App\\AppUser	Usuarios de las aplicaciones	f	2017-10-11 20:25:59	2017-10-11 20:25:59	0	Controller	\N
4	brands	brands	Brand	Brands	voyager-tag	App\\Brand	Marcas de los dispositivos	f	2017-10-11 20:25:59	2017-10-11 20:25:59	0	Controller	\N
5	browsers	browsers	Browser	Browsers	voyager-browser	App\\Browser	Navegadores	f	2017-10-11 20:25:59	2017-10-11 20:25:59	0	Controller	\N
6	devices	devices	Device	Devices	voyager-phone	App\\Device	Dispositivos móviles	f	2017-10-11 20:25:59	2017-10-11 20:25:59	0	DataTablesController	\N
7	messages	messages	Message	Messages	voyager-chat	App\\Message	Mensajes	f	2017-10-11 20:25:59	2017-10-11 20:25:59	0	MessagesController	\N
8	platforms	platforms	Platform	Platforms	voyager-laptop	App\\Platform	Plataformas donde andan las aplicaciones	f	2017-10-11 20:25:59	2017-10-11 20:25:59	0	Controller	\N
9	ranges	ranges	Range	Ranges	voyager-star-half	App\\Range	Rangos de calificaciones	f	2017-10-11 20:25:59	2017-10-11 20:25:59	0	Controller	\N
10	ratings	ratings	Rating	Ratings	voyager-star-two	App\\Rating	Calificaciones	f	2017-10-11 20:25:59	2017-10-11 20:25:59	0	DataTablesController	\N
11	menus	menus	Menu	Menus	voyager-list	TCG\\Voyager\\Models\\Menu	Menús	f	2017-10-11 20:25:59	2017-10-11 20:25:59	0		\N
12	roles	roles	Role	Roles	voyager-lock	TCG\\Voyager\\Models\\Role		f	2017-10-11 20:25:59	2017-10-11 20:25:59	0	Controller	\N
\.


--
-- Data for Name: devices; Type: TABLE DATA; Schema: public; Owner: hermes
--

COPY devices (id, name, screen_width, screen_height, ppi, brand_id, platform_id, created_at, updated_at, deleted_at) FROM stdin;
1	iPhone 6s	750	1334	326	1	1	2017-10-11 20:25:59	\N	\N
2	Galaxy S7	1440	2560	557	2	2	2017-10-11 20:25:59	\N	\N
3	Pixel XL	1440	2560	534	3	2	2017-10-11 20:25:59	\N	\N
4	iPhone 7	750	1334	326	1	1	2017-10-11 20:25:59	\N	\N
5	Galaxy J7	720	1280	267	2	2	2017-10-11 20:25:59	\N	\N
\.


--
-- Data for Name: failed_jobs; Type: TABLE DATA; Schema: public; Owner: hermes
--

COPY failed_jobs (id, connection, queue, payload, exception, failed_at) FROM stdin;
\.


--
-- Data for Name: menu_items; Type: TABLE DATA; Schema: public; Owner: hermes
--

COPY menu_items (id, menu_id, title, url, target, icon_class, color, parent_id, "order", created_at, updated_at, route, parameters) FROM stdin;
1	1	Dashboard	/admin	_self	voyager-bar-chart	\N	\N	1	2017-10-11 20:25:59	2017-10-11 20:25:59	\N	\N
2	1	Calificaciones		_self	voyager-star-two	\N	\N	2	2017-10-11 20:25:59	2017-10-11 20:25:59	voyager.ratings.index	\N
3	1	Mensajes		_self	voyager-chat	\N	\N	3	2017-10-11 20:25:59	2017-10-11 20:25:59	voyager.messages.index	\N
4	1	Usuarios		_self	voyager-people	\N	\N	4	2017-10-11 20:25:59	2017-10-11 20:25:59	voyager.appusers.index	\N
5	1	Aplicaciones		_self	voyager-categories	\N	\N	5	2017-10-11 20:25:59	2017-10-11 20:25:59	voyager.apps.index	\N
6	1	Rangos		_self	voyager-star-half	\N	\N	6	2017-10-11 20:25:59	2017-10-11 20:25:59	voyager.ranges.index	\N
7	1	Contexto		_self	voyager-world	\N	\N	7	2017-10-11 20:25:59	2017-10-11 20:25:59	\N	\N
8	1	Plataformas		_self	voyager-laptop	\N	7	1	2017-10-11 20:25:59	2017-10-11 20:25:59	voyager.platforms.index	\N
9	1	Dispositivos		_self	voyager-phone	\N	7	2	2017-10-11 20:25:59	2017-10-11 20:25:59	voyager.devices.index	\N
10	1	Marcas		_self	voyager-tag	\N	7	3	2017-10-11 20:25:59	2017-10-11 20:25:59	voyager.brands.index	\N
11	1	Navegadores		_self	voyager-browser	\N	7	4	2017-10-11 20:25:59	2017-10-11 20:25:59	voyager.browsers.index	\N
12	1	Administración		_self	voyager-settings	\N	\N	12	2017-10-11 20:25:59	2017-10-11 20:25:59	\N	\N
13	1	Personal	/admin/users	_self	voyager-person	\N	12	1	2017-10-11 20:25:59	2017-10-11 20:25:59	\N	\N
14	1	Roles	/admin/roles	_self	voyager-lock	\N	12	2	2017-10-11 20:25:59	2017-10-11 20:25:59	\N	\N
15	1	Menús	/admin/menus	_self	voyager-list	\N	12	3	2017-10-11 20:25:59	2017-10-11 20:25:59	\N	\N
16	1	DB	/admin/database	_self	voyager-data	\N	12	4	2017-10-11 20:25:59	2017-10-11 20:25:59	\N	\N
17	1	Configuración	/admin/settings	_self	voyager-tools	\N	12	5	2017-10-11 20:25:59	2017-10-11 20:25:59	\N	\N
\.


--
-- Data for Name: menus; Type: TABLE DATA; Schema: public; Owner: hermes
--

COPY menus (id, name, created_at, updated_at) FROM stdin;
1	admin	2017-10-11 20:25:59	2017-10-11 20:25:59
\.


--
-- Data for Name: messages; Type: TABLE DATA; Schema: public; Owner: hermes
--

COPY messages (id, message, direction, status, transport_id, rating_id, created_at, updated_at) FROM stdin;
1	Lorem ipsum dolor sit amet, consectetur adipiscing elit sed eiusmod tempor incidunt ut labore et dolore magna aliqua.	in	0		1	2017-10-11 20:26:00	\N
2	Lorem ipsum dolor sit amet, consectetur adipiscing elit sed eiusmod tempor incidunt ut labore et dolore magna aliqua.	in	0		2	2017-10-11 20:26:00	\N
3	Lorem ipsum dolor sit amet, consectetur adipiscing elit sed eiusmod tempor incidunt ut labore et dolore magna aliqua.	in	0		4	2017-10-11 20:26:00	\N
4	Lorem ipsum dolor sit amet, consectetur adipiscing elit sed eiusmod tempor incidunt ut labore et dolore magna aliqua.	in	0		5	2017-10-11 20:26:00	\N
5	Lorem ipsum dolor sit amet, consectetur adipiscing elit sed eiusmod tempor incidunt ut labore et dolore magna aliqua.	in	0		6	2017-10-11 20:26:00	\N
6	Lorem ipsum dolor sit amet, consectetur adipiscing elit sed eiusmod tempor incidunt ut labore et dolore magna aliqua.	in	0		8	2017-10-11 20:26:00	\N
7	Lorem ipsum dolor sit amet, consectetur adipiscing elit sed eiusmod tempor incidunt ut labore et dolore magna aliqua.	in	0		9	2017-10-11 20:26:00	\N
8	Lorem ipsum dolor sit amet, consectetur adipiscing elit sed eiusmod tempor incidunt ut labore et dolore magna aliqua.	in	0		10	2017-10-11 20:26:00	\N
9	Lorem ipsum dolor sit amet, consectetur adipiscing elit sed eiusmod tempor incidunt ut labore et dolore magna aliqua.	in	0		11	2017-10-11 20:26:00	\N
10	Lorem ipsum dolor sit amet, consectetur adipiscing elit sed eiusmod tempor incidunt ut labore et dolore magna aliqua.	in	0		14	2017-10-11 20:26:00	\N
11	Lorem ipsum dolor sit amet, consectetur adipiscing elit sed eiusmod tempor incidunt ut labore et dolore magna aliqua.	in	0		15	2017-10-11 20:26:00	\N
12	Lorem ipsum dolor sit amet, consectetur adipiscing elit sed eiusmod tempor incidunt ut labore et dolore magna aliqua.	in	0		16	2017-10-11 20:26:00	\N
13	Lorem ipsum dolor sit amet, consectetur adipiscing elit sed eiusmod tempor incidunt ut labore et dolore magna aliqua.	in	0		17	2017-10-11 20:26:00	\N
\.


--
-- Data for Name: migrations; Type: TABLE DATA; Schema: public; Owner: hermes
--

COPY migrations (id, migration, batch) FROM stdin;
1215	2014_10_12_000000_create_users_table	1
1216	2014_10_12_100000_create_password_resets_table	1
1217	2016_01_01_000000_add_voyager_user_fields	1
1218	2016_01_01_000000_create_data_types_table	1
1219	2016_01_01_000000_create_pages_table	1
1220	2016_01_01_000000_create_posts_table	1
1221	2016_02_15_204651_create_categories_table	1
1222	2016_05_19_173453_create_menu_table	1
1223	2016_10_21_190000_create_roles_table	1
1224	2016_10_21_190000_create_settings_table	1
1225	2016_11_30_135954_create_permission_table	1
1226	2016_11_30_141208_create_permission_role_table	1
1227	2016_12_26_201236_data_types__add__server_side	1
1228	2017_01_13_000000_add_route_to_menu_items_table	1
1229	2017_01_14_005015_create_translations_table	1
1230	2017_01_15_000000_add_permission_group_id_to_permissions_table	1
1231	2017_01_15_000000_create_permission_groups_table	1
1232	2017_01_15_000000_make_table_name_nullable_in_permissions_table	1
1233	2017_03_06_000000_add_controller_to_data_types_table	1
1234	2017_04_11_000000_alter_post_nullable_fields_table	1
1235	2017_04_21_000000_add_order_to_data_rows_table	1
1236	2017_05_30_140334_create_app_users_table	1
1237	2017_05_30_144303_create_apps_table	1
1238	2017_05_30_144304_create_ranges_table	1
1239	2017_05_30_144305_create_platforms_table	1
1240	2017_05_30_144306_create_brands_table	1
1241	2017_05_30_144307_create_devices_table	1
1242	2017_05_30_144308_create_browsers_table	1
1243	2017_05_30_144309_create_ratings_table	1
1244	2017_06_01_140433_create_messages_table	1
1245	2017_06_07_224132_create_app_platform_table	1
1246	2017_06_07_232505_create_app_user_platform_table	1
1247	2017_06_08_132835_create_app_user_table	1
1248	2017_06_08_140926_create_app_user_app_table	1
1249	2017_06_13_202957_create_app_user_device_table	1
1250	2017_07_05_210000_add_policyname_to_data_types_table	1
1251	2017_08_05_000000_add_group_to_settings_table	1
1252	2017_10_10_220826_create_failed_jobs_table	1
\.


--
-- Data for Name: pages; Type: TABLE DATA; Schema: public; Owner: hermes
--

COPY pages (id, author_id, title, excerpt, body, image, slug, meta_description, meta_keywords, status, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: password_resets; Type: TABLE DATA; Schema: public; Owner: hermes
--

COPY password_resets (email, token, created_at) FROM stdin;
\.


--
-- Data for Name: permission_groups; Type: TABLE DATA; Schema: public; Owner: hermes
--

COPY permission_groups (id, name) FROM stdin;
\.


--
-- Data for Name: permission_role; Type: TABLE DATA; Schema: public; Owner: hermes
--

COPY permission_role (permission_id, role_id) FROM stdin;
1	1
2	1
3	1
4	1
5	1
6	1
7	1
8	1
9	1
10	1
11	1
12	1
13	1
14	1
15	1
16	1
17	1
18	1
19	1
20	1
21	1
22	1
23	1
24	1
25	1
26	1
27	1
28	1
29	1
32	1
33	1
34	1
35	1
36	1
37	1
38	1
39	1
40	1
41	1
42	1
43	1
44	1
47	1
48	1
49	1
52	1
53	1
54	1
57	1
58	1
59	1
61	1
62	1
63	1
64	1
66	1
67	1
68	1
69	1
72	1
1	2
2	2
4	2
5	2
6	2
7	2
8	2
9	2
10	2
11	2
12	2
23	2
24	2
25	2
26	2
27	2
28	2
29	2
33	2
34	2
35	2
36	2
37	2
38	2
39	2
40	2
41	2
43	2
44	2
48	2
49	2
53	2
54	2
58	2
59	2
61	2
63	2
64	2
66	2
67	2
68	2
69	2
1	3
2	3
4	3
5	3
6	3
7	3
8	3
9	3
10	3
11	3
12	3
28	3
29	3
33	3
34	3
38	3
39	3
43	3
44	3
48	3
49	3
53	3
54	3
58	3
59	3
63	3
64	3
66	3
67	3
68	3
69	3
1	4
2	4
4	4
5	4
6	4
7	4
8	4
9	4
10	4
11	4
12	4
28	4
29	4
33	4
34	4
38	4
39	4
43	4
44	4
48	4
49	4
53	4
54	4
58	4
59	4
63	4
64	4
68	4
69	4
\.


--
-- Data for Name: permissions; Type: TABLE DATA; Schema: public; Owner: hermes
--

COPY permissions (id, key, table_name, created_at, updated_at, permission_group_id) FROM stdin;
1	browse_admin	\N	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
2	browse_database	\N	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
3	browse_settings	\N	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
4	browse_apps	\N	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
5	browse_appusers	\N	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
6	browse_platforms	\N	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
7	browse_brands	\N	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
8	browse_browsers	\N	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
9	browse_devices	\N	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
10	browse_ranges	\N	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
11	browse_messages	\N	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
12	browse_ratings	\N	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
13	browse_menus	menus	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
14	read_menus	menus	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
15	edit_menus	menus	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
16	add_menus	menus	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
17	delete_menus	menus	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
18	browse_roles	roles	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
19	read_roles	roles	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
20	edit_roles	roles	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
21	add_roles	roles	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
22	delete_roles	roles	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
23	browse_users	users	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
24	read_users	users	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
25	edit_users	users	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
26	add_users	users	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
27	delete_users	users	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
28	browse_appusers	appusers	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
29	read_appusers	appusers	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
30	edit_appusers	appusers	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
31	add_appusers	appusers	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
32	delete_appusers	appusers	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
33	browse_apps	apps	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
34	read_apps	apps	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
35	edit_apps	apps	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
36	add_apps	apps	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
37	delete_apps	apps	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
38	browse_platforms	platforms	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
39	read_platforms	platforms	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
40	edit_platforms	platforms	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
41	add_platforms	platforms	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
42	delete_platforms	platforms	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
43	browse_brands	brands	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
44	read_brands	brands	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
45	edit_brands	brands	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
46	add_brands	brands	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
47	delete_brands	brands	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
48	browse_browsers	browsers	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
49	read_browsers	browsers	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
50	edit_browsers	browsers	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
51	add_browsers	browsers	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
52	delete_browsers	browsers	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
53	browse_devices	devices	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
54	read_devices	devices	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
55	edit_devices	devices	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
56	add_devices	devices	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
57	delete_devices	devices	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
58	browse_ranges	ranges	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
59	read_ranges	ranges	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
60	edit_ranges	ranges	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
61	add_ranges	ranges	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
62	delete_ranges	ranges	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
63	browse_messages	messages	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
64	read_messages	messages	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
65	edit_messages	messages	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
66	add_messages	messages	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
67	delete_messages	messages	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
68	browse_ratings	ratings	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
69	read_ratings	ratings	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
70	edit_ratings	ratings	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
71	add_ratings	ratings	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
72	delete_ratings	ratings	2017-10-11 20:25:59	2017-10-11 20:25:59	\N
\.


--
-- Data for Name: platforms; Type: TABLE DATA; Schema: public; Owner: hermes
--

COPY platforms (id, name, key, created_at, updated_at, deleted_at) FROM stdin;
1	iOS	e10adc3949ba59abbe56e057f20f883e	2017-10-11 20:25:59	\N	\N
2	Android	c33367701511b4f6020ec61ded352059	2017-10-11 20:25:59	\N	\N
3	Browser	0d06fd8cb26eb57f1a690f493663cc55	2017-10-11 20:25:59	\N	\N
\.


--
-- Data for Name: posts; Type: TABLE DATA; Schema: public; Owner: hermes
--

COPY posts (id, author_id, category_id, title, seo_title, excerpt, body, image, slug, meta_description, meta_keywords, status, featured, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: ranges; Type: TABLE DATA; Schema: public; Owner: hermes
--

COPY ranges (id, name, "from", "to", key, created_at, updated_at, deleted_at) FROM stdin;
1	0/5	0	5	e10adc3949ba59abbe56e057f20f883e	2017-10-11 20:25:59	\N	\N
2	0/10	0	10	c33367701511b4f6020ec61ded352059	2017-10-11 20:25:59	\N	\N
3	-5/5	-5	5	6c44e5cd17f0019c64b042e4a745412a	2017-10-11 20:25:59	\N	\N
\.


--
-- Data for Name: ratings; Type: TABLE DATA; Schema: public; Owner: hermes
--

COPY ratings (id, rating, description, app_version, browser_version, platform_version, has_message, app_id, range_id, platform_id, device_id, appuser_id, browser_id, created_at, updated_at, deleted_at) FROM stdin;
1	5	Muy bueno	2.0	\N	9.0	t	1	1	1	1	1	\N	2017-10-11 20:25:59	\N	\N
2	2	Malo	2.0	\N	8.0	t	1	1	1	4	2	\N	2017-10-11 20:25:59	\N	\N
3	1	Muy malo	1.0	\N	8.0	f	1	2	1	4	3	\N	2017-10-11 20:26:00	\N	\N
4	3	Regular	1.0	\N	6.0	t	1	3	2	5	4	\N	2017-10-11 20:26:00	\N	\N
5	4	Bueno	2.0	\N	5.1	t	1	1	2	2	5	\N	2017-10-11 20:26:00	\N	\N
6	3	Regular	1.0	\N	8.0	t	1	2	1	4	6	\N	2017-10-11 20:26:00	\N	\N
7	2	Malo	1.0	\N	6.0	f	2	1	2	2	1	\N	2017-10-11 20:26:00	\N	\N
8	3	Regular	2.0	\N	6.0	t	2	3	2	2	2	\N	2017-10-11 20:26:00	\N	\N
9	4	Bueno	2.0	\N	9.0	t	2	2	1	4	3	\N	2017-10-11 20:26:00	\N	\N
10	4	Bueno	2.0	\N	9.0	t	2	3	1	1	4	\N	2017-10-11 20:26:00	\N	\N
11	2	Malo	1.0	\N	9.0	t	2	2	1	1	5	\N	2017-10-11 20:26:00	\N	\N
12	3	Regular	1.0	\N	8.0	f	2	1	1	1	6	\N	2017-10-11 20:26:00	\N	\N
13	4	Bueno	2.0	\N	6.0	f	3	2	2	2	1	\N	2017-10-11 20:26:00	\N	\N
14	4	Bueno	2.0	\N	5.1	t	3	2	2	5	2	\N	2017-10-11 20:26:00	\N	\N
15	3	Regular	2.0	\N	5.1	t	3	2	2	5	3	\N	2017-10-11 20:26:00	\N	\N
16	5	Muy bueno	2.0	\N	6.0	t	3	3	2	2	4	\N	2017-10-11 20:26:00	\N	\N
17	2	Malo	1.0	\N	8.0	t	3	2	1	1	5	\N	2017-10-11 20:26:00	\N	\N
18	4	Bueno	1.0	\N	9.0	f	3	3	1	4	6	\N	2017-10-11 20:26:00	\N	\N
\.


--
-- Data for Name: roles; Type: TABLE DATA; Schema: public; Owner: hermes
--

COPY roles (id, name, display_name, created_at, updated_at) FROM stdin;
1	admin	Administrador	2017-10-11 20:25:59	2017-10-11 20:25:59
2	supervisor	Supervisor	2017-10-11 20:25:59	2017-10-11 20:25:59
3	support	Soporte	2017-10-11 20:25:59	2017-10-11 20:25:59
4	user	Usuario	2017-10-11 20:25:59	2017-10-11 20:25:59
\.


--
-- Data for Name: settings; Type: TABLE DATA; Schema: public; Owner: hermes
--

COPY settings (id, key, display_name, value, details, type, "order", "group") FROM stdin;
1	title	Site Title			text	1	\N
2	description	Site Description			text	2	\N
3	logo	Site Logo			image	3	\N
4	admin_bg_image	Admin Background Image			image	9	\N
5	admin_title	Admin Title	Hermes		text	4	\N
6	admin_description	Admin Description	Gestión de feedback de las apps de la Ciudad		text	5	\N
7	admin_loader	Admin Loader			image	6	\N
8	admin_icon_image	Admin Icon Image			image	7	\N
9	google_analytics_client_id	Google Analytics Client ID			text	9	\N
\.


--
-- Data for Name: translations; Type: TABLE DATA; Schema: public; Owner: hermes
--

COPY translations (id, table_name, column_name, foreign_key, locale, value, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: hermes
--

COPY users (id, name, email, remember_token, created_at, updated_at, updated_by, avatar, role_id) FROM stdin;
1	Admin	admin@admin.com	0prXpLH313SzHwgiNHAaGjMwBELfMb7GORtwDrbpjvxkiowrMbkBzGQfB7PE	2017-10-11 20:25:59	\N	\N	users/default.png	1
2	Juan Fernández	juan@fernandez.com	zUivg2tOekAChn5CQY1Ltv8hDAWGWHkkwpVyKBnWBdc25y6CbtEMN3OTitJP	2017-10-11 20:25:59	\N	\N	users/default.png	3
3	Martina Giménez	martina@gimenez.com	HrpBsfe2AYjD388l1QYLE6BRdNheQRyfV85HeSo28Qqb602i5YgobH1o2dqU	2017-10-11 20:25:59	\N	\N	users/default.png	3
4	Paula Carrizo	paula@carrizo.com	VefsJvzrviENP4SchId8p5lymOE0nfAqZCeCaHwdy6CeFE1vTlXIqJUl04Yi	2017-10-11 20:25:59	\N	\N	users/default.png	3
5	Miguel Rodríguez	miguel@rodriguez.com	1czFDvSXyrnavkiXH9VxH4cSIApvWxkdWTk0v2I0lsVt6bn200ghQRBD9Z0w	2017-10-11 20:25:59	\N	\N	users/default.png	4
6	Sofía Estévez	sofia@estevez.com	NeWeO2lJoIa7i8TtZ5idcC23XNZnoHfEeN3SYkBvQGGhjRbdMH3CLrprqjvw	2017-10-11 20:25:59	\N	\N	users/default.png	4
7	Nicolás Uriarte	nicolas@uriarte.com	cjOzlhH6IXmURo2qGbxfBJtwLf9Hdl8K5SB32WnSeW73J9DPTGIJKZ2jwT0N	2017-10-11 20:25:59	\N	\N	users/default.png	4
\.


--
-- Name: apps_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hermes
--

SELECT pg_catalog.setval('apps_id_seq', 3, true);


--
-- Name: appusers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hermes
--

SELECT pg_catalog.setval('appusers_id_seq', 6, true);


--
-- Name: brands_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hermes
--

SELECT pg_catalog.setval('brands_id_seq', 3, true);


--
-- Name: browsers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hermes
--

SELECT pg_catalog.setval('browsers_id_seq', 3, true);


--
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hermes
--

SELECT pg_catalog.setval('categories_id_seq', 1, false);


--
-- Name: data_rows_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hermes
--

SELECT pg_catalog.setval('data_rows_id_seq', 96, true);


--
-- Name: data_types_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hermes
--

SELECT pg_catalog.setval('data_types_id_seq', 12, true);


--
-- Name: devices_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hermes
--

SELECT pg_catalog.setval('devices_id_seq', 5, true);


--
-- Name: failed_jobs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hermes
--

SELECT pg_catalog.setval('failed_jobs_id_seq', 1, false);


--
-- Name: menu_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hermes
--

SELECT pg_catalog.setval('menu_items_id_seq', 17, true);


--
-- Name: menus_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hermes
--

SELECT pg_catalog.setval('menus_id_seq', 1, true);


--
-- Name: messages_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hermes
--

SELECT pg_catalog.setval('messages_id_seq', 13, true);


--
-- Name: migrations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hermes
--

SELECT pg_catalog.setval('migrations_id_seq', 1252, true);


--
-- Name: pages_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hermes
--

SELECT pg_catalog.setval('pages_id_seq', 1, false);


--
-- Name: permission_groups_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hermes
--

SELECT pg_catalog.setval('permission_groups_id_seq', 1, false);


--
-- Name: permissions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hermes
--

SELECT pg_catalog.setval('permissions_id_seq', 72, true);


--
-- Name: platforms_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hermes
--

SELECT pg_catalog.setval('platforms_id_seq', 3, true);


--
-- Name: posts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hermes
--

SELECT pg_catalog.setval('posts_id_seq', 1, false);


--
-- Name: ranges_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hermes
--

SELECT pg_catalog.setval('ranges_id_seq', 3, true);


--
-- Name: ratings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hermes
--

SELECT pg_catalog.setval('ratings_id_seq', 18, true);


--
-- Name: roles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hermes
--

SELECT pg_catalog.setval('roles_id_seq', 4, true);


--
-- Name: settings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hermes
--

SELECT pg_catalog.setval('settings_id_seq', 9, true);


--
-- Name: translations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hermes
--

SELECT pg_catalog.setval('translations_id_seq', 1, false);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: hermes
--

SELECT pg_catalog.setval('users_id_seq', 7, true);


--
-- Name: app_platform app_platform_pkey; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY app_platform
    ADD CONSTRAINT app_platform_pkey PRIMARY KEY (platform_id, app_id);


--
-- Name: app_user_app app_user_app_pkey; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY app_user_app
    ADD CONSTRAINT app_user_app_pkey PRIMARY KEY (app_id, app_user_id);


--
-- Name: app_user_device app_user_device_pkey; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY app_user_device
    ADD CONSTRAINT app_user_device_pkey PRIMARY KEY (device_id, app_user_id);


--
-- Name: app_user app_user_pkey; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY app_user
    ADD CONSTRAINT app_user_pkey PRIMARY KEY (app_id, user_id);


--
-- Name: app_user_platform app_user_platform_pkey; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY app_user_platform
    ADD CONSTRAINT app_user_platform_pkey PRIMARY KEY (platform_id, app_user_id);


--
-- Name: apps apps_key_unique; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY apps
    ADD CONSTRAINT apps_key_unique UNIQUE (key);


--
-- Name: apps apps_name_unique; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY apps
    ADD CONSTRAINT apps_name_unique UNIQUE (name);


--
-- Name: apps apps_pkey; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY apps
    ADD CONSTRAINT apps_pkey PRIMARY KEY (id);


--
-- Name: appusers appusers_pkey; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY appusers
    ADD CONSTRAINT appusers_pkey PRIMARY KEY (id);


--
-- Name: brands brands_name_unique; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY brands
    ADD CONSTRAINT brands_name_unique UNIQUE (name);


--
-- Name: brands brands_pkey; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY brands
    ADD CONSTRAINT brands_pkey PRIMARY KEY (id);


--
-- Name: browsers browsers_name_unique; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY browsers
    ADD CONSTRAINT browsers_name_unique UNIQUE (name);


--
-- Name: browsers browsers_pkey; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY browsers
    ADD CONSTRAINT browsers_pkey PRIMARY KEY (id);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- Name: categories categories_slug_unique; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY categories
    ADD CONSTRAINT categories_slug_unique UNIQUE (slug);


--
-- Name: data_rows data_rows_pkey; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY data_rows
    ADD CONSTRAINT data_rows_pkey PRIMARY KEY (id);


--
-- Name: data_types data_types_name_unique; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY data_types
    ADD CONSTRAINT data_types_name_unique UNIQUE (name);


--
-- Name: data_types data_types_pkey; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY data_types
    ADD CONSTRAINT data_types_pkey PRIMARY KEY (id);


--
-- Name: data_types data_types_slug_unique; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY data_types
    ADD CONSTRAINT data_types_slug_unique UNIQUE (slug);


--
-- Name: devices devices_name_unique; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY devices
    ADD CONSTRAINT devices_name_unique UNIQUE (name);


--
-- Name: devices devices_pkey; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY devices
    ADD CONSTRAINT devices_pkey PRIMARY KEY (id);


--
-- Name: failed_jobs failed_jobs_pkey; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY failed_jobs
    ADD CONSTRAINT failed_jobs_pkey PRIMARY KEY (id);


--
-- Name: menu_items menu_items_pkey; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY menu_items
    ADD CONSTRAINT menu_items_pkey PRIMARY KEY (id);


--
-- Name: menus menus_name_unique; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY menus
    ADD CONSTRAINT menus_name_unique UNIQUE (name);


--
-- Name: menus menus_pkey; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY menus
    ADD CONSTRAINT menus_pkey PRIMARY KEY (id);


--
-- Name: messages messages_pkey; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY messages
    ADD CONSTRAINT messages_pkey PRIMARY KEY (id);


--
-- Name: migrations migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY migrations
    ADD CONSTRAINT migrations_pkey PRIMARY KEY (id);


--
-- Name: pages pages_pkey; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY pages
    ADD CONSTRAINT pages_pkey PRIMARY KEY (id);


--
-- Name: pages pages_slug_unique; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY pages
    ADD CONSTRAINT pages_slug_unique UNIQUE (slug);


--
-- Name: permission_groups permission_groups_name_unique; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY permission_groups
    ADD CONSTRAINT permission_groups_name_unique UNIQUE (name);


--
-- Name: permission_groups permission_groups_pkey; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY permission_groups
    ADD CONSTRAINT permission_groups_pkey PRIMARY KEY (id);


--
-- Name: permission_role permission_role_pkey; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY permission_role
    ADD CONSTRAINT permission_role_pkey PRIMARY KEY (permission_id, role_id);


--
-- Name: permissions permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY permissions
    ADD CONSTRAINT permissions_pkey PRIMARY KEY (id);


--
-- Name: platforms platforms_key_unique; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY platforms
    ADD CONSTRAINT platforms_key_unique UNIQUE (key);


--
-- Name: platforms platforms_name_unique; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY platforms
    ADD CONSTRAINT platforms_name_unique UNIQUE (name);


--
-- Name: platforms platforms_pkey; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY platforms
    ADD CONSTRAINT platforms_pkey PRIMARY KEY (id);


--
-- Name: posts posts_pkey; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY posts
    ADD CONSTRAINT posts_pkey PRIMARY KEY (id);


--
-- Name: posts posts_slug_unique; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY posts
    ADD CONSTRAINT posts_slug_unique UNIQUE (slug);


--
-- Name: ranges ranges_key_unique; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY ranges
    ADD CONSTRAINT ranges_key_unique UNIQUE (key);


--
-- Name: ranges ranges_pkey; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY ranges
    ADD CONSTRAINT ranges_pkey PRIMARY KEY (id);


--
-- Name: ratings ratings_pkey; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY ratings
    ADD CONSTRAINT ratings_pkey PRIMARY KEY (id);


--
-- Name: roles roles_name_unique; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY roles
    ADD CONSTRAINT roles_name_unique UNIQUE (name);


--
-- Name: roles roles_pkey; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);


--
-- Name: settings settings_key_unique; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY settings
    ADD CONSTRAINT settings_key_unique UNIQUE (key);


--
-- Name: settings settings_pkey; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY settings
    ADD CONSTRAINT settings_pkey PRIMARY KEY (id);


--
-- Name: translations translations_pkey; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY translations
    ADD CONSTRAINT translations_pkey PRIMARY KEY (id);


--
-- Name: translations translations_table_name_column_name_foreign_key_locale_unique; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY translations
    ADD CONSTRAINT translations_table_name_column_name_foreign_key_locale_unique UNIQUE (table_name, column_name, foreign_key, locale);


--
-- Name: users users_email_unique; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY users
    ADD CONSTRAINT users_email_unique UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: app_platform_app_id_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX app_platform_app_id_index ON app_platform USING btree (app_id);


--
-- Name: app_platform_platform_id_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX app_platform_platform_id_index ON app_platform USING btree (platform_id);


--
-- Name: app_user_app_app_id_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX app_user_app_app_id_index ON app_user_app USING btree (app_id);


--
-- Name: app_user_app_app_user_id_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX app_user_app_app_user_id_index ON app_user_app USING btree (app_user_id);


--
-- Name: app_user_app_id_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX app_user_app_id_index ON app_user USING btree (app_id);


--
-- Name: app_user_device_app_user_id_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX app_user_device_app_user_id_index ON app_user_device USING btree (app_user_id);


--
-- Name: app_user_device_device_id_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX app_user_device_device_id_index ON app_user_device USING btree (device_id);


--
-- Name: app_user_platform_app_user_id_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX app_user_platform_app_user_id_index ON app_user_platform USING btree (app_user_id);


--
-- Name: app_user_platform_platform_id_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX app_user_platform_platform_id_index ON app_user_platform USING btree (platform_id);


--
-- Name: app_user_user_id_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX app_user_user_id_index ON app_user USING btree (user_id);


--
-- Name: apps_key_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX apps_key_index ON apps USING btree (key);


--
-- Name: apps_name_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX apps_name_index ON apps USING btree (name);


--
-- Name: appusers_email_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX appusers_email_index ON appusers USING btree (email);


--
-- Name: appusers_miba_id_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX appusers_miba_id_index ON appusers USING btree (miba_id);


--
-- Name: devices_brand_id_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX devices_brand_id_index ON devices USING btree (brand_id);


--
-- Name: devices_name_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX devices_name_index ON devices USING btree (name);


--
-- Name: devices_platform_id_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX devices_platform_id_index ON devices USING btree (platform_id);


--
-- Name: messages_direction_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX messages_direction_index ON messages USING btree (direction);


--
-- Name: messages_rating_id_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX messages_rating_id_index ON messages USING btree (rating_id);


--
-- Name: messages_status_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX messages_status_index ON messages USING btree (status);


--
-- Name: messages_transport_id_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX messages_transport_id_index ON messages USING btree (transport_id);


--
-- Name: password_resets_email_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX password_resets_email_index ON password_resets USING btree (email);


--
-- Name: permission_role_permission_id_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX permission_role_permission_id_index ON permission_role USING btree (permission_id);


--
-- Name: permission_role_role_id_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX permission_role_role_id_index ON permission_role USING btree (role_id);


--
-- Name: permissions_key_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX permissions_key_index ON permissions USING btree (key);


--
-- Name: platforms_key_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX platforms_key_index ON platforms USING btree (key);


--
-- Name: ranges_key_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX ranges_key_index ON ranges USING btree (key);


--
-- Name: ratings_app_id_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX ratings_app_id_index ON ratings USING btree (app_id);


--
-- Name: ratings_appuser_id_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX ratings_appuser_id_index ON ratings USING btree (appuser_id);


--
-- Name: ratings_browser_id_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX ratings_browser_id_index ON ratings USING btree (browser_id);


--
-- Name: ratings_device_id_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX ratings_device_id_index ON ratings USING btree (device_id);


--
-- Name: ratings_has_message_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX ratings_has_message_index ON ratings USING btree (has_message);


--
-- Name: ratings_platform_id_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX ratings_platform_id_index ON ratings USING btree (platform_id);


--
-- Name: ratings_range_id_index; Type: INDEX; Schema: public; Owner: hermes
--

CREATE INDEX ratings_range_id_index ON ratings USING btree (range_id);


--
-- Name: app_platform app_platform_app_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY app_platform
    ADD CONSTRAINT app_platform_app_id_foreign FOREIGN KEY (app_id) REFERENCES apps(id) ON DELETE CASCADE;


--
-- Name: app_platform app_platform_platform_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY app_platform
    ADD CONSTRAINT app_platform_platform_id_foreign FOREIGN KEY (platform_id) REFERENCES platforms(id) ON DELETE CASCADE;


--
-- Name: app_user_app app_user_app_app_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY app_user_app
    ADD CONSTRAINT app_user_app_app_id_foreign FOREIGN KEY (app_id) REFERENCES apps(id) ON DELETE CASCADE;


--
-- Name: app_user_app app_user_app_app_user_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY app_user_app
    ADD CONSTRAINT app_user_app_app_user_id_foreign FOREIGN KEY (app_user_id) REFERENCES appusers(id) ON DELETE CASCADE;


--
-- Name: app_user app_user_app_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY app_user
    ADD CONSTRAINT app_user_app_id_foreign FOREIGN KEY (app_id) REFERENCES apps(id) ON DELETE CASCADE;


--
-- Name: app_user_device app_user_device_app_user_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY app_user_device
    ADD CONSTRAINT app_user_device_app_user_id_foreign FOREIGN KEY (app_user_id) REFERENCES appusers(id) ON DELETE CASCADE;


--
-- Name: app_user_device app_user_device_device_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY app_user_device
    ADD CONSTRAINT app_user_device_device_id_foreign FOREIGN KEY (device_id) REFERENCES devices(id) ON DELETE CASCADE;


--
-- Name: app_user_platform app_user_platform_app_user_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY app_user_platform
    ADD CONSTRAINT app_user_platform_app_user_id_foreign FOREIGN KEY (app_user_id) REFERENCES appusers(id) ON DELETE CASCADE;


--
-- Name: app_user_platform app_user_platform_platform_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY app_user_platform
    ADD CONSTRAINT app_user_platform_platform_id_foreign FOREIGN KEY (platform_id) REFERENCES platforms(id) ON DELETE CASCADE;


--
-- Name: app_user app_user_user_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY app_user
    ADD CONSTRAINT app_user_user_id_foreign FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;


--
-- Name: apps apps_updated_by_foreign; Type: FK CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY apps
    ADD CONSTRAINT apps_updated_by_foreign FOREIGN KEY (updated_by) REFERENCES users(id) ON DELETE SET NULL;


--
-- Name: categories categories_parent_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY categories
    ADD CONSTRAINT categories_parent_id_foreign FOREIGN KEY (parent_id) REFERENCES categories(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- Name: data_rows data_rows_data_type_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY data_rows
    ADD CONSTRAINT data_rows_data_type_id_foreign FOREIGN KEY (data_type_id) REFERENCES data_types(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: devices devices_brand_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY devices
    ADD CONSTRAINT devices_brand_id_foreign FOREIGN KEY (brand_id) REFERENCES brands(id) ON DELETE SET NULL;


--
-- Name: devices devices_platform_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY devices
    ADD CONSTRAINT devices_platform_id_foreign FOREIGN KEY (platform_id) REFERENCES platforms(id) ON DELETE SET NULL;


--
-- Name: menu_items menu_items_menu_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY menu_items
    ADD CONSTRAINT menu_items_menu_id_foreign FOREIGN KEY (menu_id) REFERENCES menus(id) ON DELETE CASCADE;


--
-- Name: messages messages_rating_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY messages
    ADD CONSTRAINT messages_rating_id_foreign FOREIGN KEY (rating_id) REFERENCES ratings(id) ON DELETE CASCADE;


--
-- Name: permission_role permission_role_permission_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY permission_role
    ADD CONSTRAINT permission_role_permission_id_foreign FOREIGN KEY (permission_id) REFERENCES permissions(id) ON DELETE CASCADE;


--
-- Name: permission_role permission_role_role_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY permission_role
    ADD CONSTRAINT permission_role_role_id_foreign FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE;


--
-- Name: ratings ratings_app_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY ratings
    ADD CONSTRAINT ratings_app_id_foreign FOREIGN KEY (app_id) REFERENCES apps(id) ON DELETE CASCADE;


--
-- Name: ratings ratings_appuser_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY ratings
    ADD CONSTRAINT ratings_appuser_id_foreign FOREIGN KEY (appuser_id) REFERENCES appusers(id) ON DELETE CASCADE;


--
-- Name: ratings ratings_browser_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY ratings
    ADD CONSTRAINT ratings_browser_id_foreign FOREIGN KEY (browser_id) REFERENCES browsers(id) ON DELETE SET NULL;


--
-- Name: ratings ratings_device_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY ratings
    ADD CONSTRAINT ratings_device_id_foreign FOREIGN KEY (device_id) REFERENCES devices(id) ON DELETE SET NULL;


--
-- Name: ratings ratings_platform_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY ratings
    ADD CONSTRAINT ratings_platform_id_foreign FOREIGN KEY (platform_id) REFERENCES platforms(id) ON DELETE SET NULL;


--
-- Name: ratings ratings_range_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY ratings
    ADD CONSTRAINT ratings_range_id_foreign FOREIGN KEY (range_id) REFERENCES ranges(id) ON DELETE CASCADE;


--
-- Name: users users_updated_by_foreign; Type: FK CONSTRAINT; Schema: public; Owner: hermes
--

ALTER TABLE ONLY users
    ADD CONSTRAINT users_updated_by_foreign FOREIGN KEY (updated_by) REFERENCES users(id) ON DELETE SET NULL;


--
-- PostgreSQL database dump complete
--

