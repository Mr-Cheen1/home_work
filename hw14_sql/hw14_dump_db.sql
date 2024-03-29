--
-- PostgreSQL database dump
--

-- Dumped from database version 16.2
-- Dumped by pg_dump version 16.2

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: orderproducts; Type: TABLE; Schema: public; Owner: store_user
--

CREATE TABLE public.orderproducts (
    order_id integer NOT NULL,
    product_id integer NOT NULL,
    quantity integer NOT NULL
);


ALTER TABLE public.orderproducts OWNER TO store_user;

--
-- Name: orders; Type: TABLE; Schema: public; Owner: store_user
--

CREATE TABLE public.orders (
    id integer NOT NULL,
    user_id integer NOT NULL,
    order_date date NOT NULL,
    total_amount numeric NOT NULL
);


ALTER TABLE public.orders OWNER TO store_user;

--
-- Name: orders_id_seq; Type: SEQUENCE; Schema: public; Owner: store_user
--

CREATE SEQUENCE public.orders_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.orders_id_seq OWNER TO store_user;

--
-- Name: orders_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: store_user
--

ALTER SEQUENCE public.orders_id_seq OWNED BY public.orders.id;


--
-- Name: products; Type: TABLE; Schema: public; Owner: store_user
--

CREATE TABLE public.products (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    price numeric NOT NULL
);


ALTER TABLE public.products OWNER TO store_user;

--
-- Name: products_id_seq; Type: SEQUENCE; Schema: public; Owner: store_user
--

CREATE SEQUENCE public.products_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.products_id_seq OWNER TO store_user;

--
-- Name: products_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: store_user
--

ALTER SEQUENCE public.products_id_seq OWNED BY public.products.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: store_user
--

CREATE TABLE public.users (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(255) NOT NULL
);


ALTER TABLE public.users OWNER TO store_user;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: store_user
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO store_user;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: store_user
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: orders id; Type: DEFAULT; Schema: public; Owner: store_user
--

ALTER TABLE ONLY public.orders ALTER COLUMN id SET DEFAULT nextval('public.orders_id_seq'::regclass);


--
-- Name: products id; Type: DEFAULT; Schema: public; Owner: store_user
--

ALTER TABLE ONLY public.products ALTER COLUMN id SET DEFAULT nextval('public.products_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: store_user
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: orderproducts; Type: TABLE DATA; Schema: public; Owner: store_user
--

COPY public.orderproducts (order_id, product_id, quantity) FROM stdin;
1	1	2
2	2	1
5	5	2
\.


--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: store_user
--

COPY public.orders (id, user_id, order_date, total_amount) FROM stdin;
5	5	2024-04-05	9000.00
1	1	2024-04-01	32000.00
2	2	2024-04-06	30000.00
\.


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: store_user
--

COPY public.products (id, name, price) FROM stdin;
1	E-book	16000.00
2	Premium coffee machine	30000.00
3	Robot Vacuum Cleaner Smart	21000.00
4	Smartphone XL	57000.00
5	Advanced fitness bracelet	9000.00
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: store_user
--

COPY public.users (id, name, email, password) FROM stdin;
1	Elena Belova	elena.newemail@example.com	newElenaPass123
2	Sergey Novikov	sergey.novikov@example.com	sergeyPass456
3	Irina Zhukova	irina.zhukova@example.com	newIrinaPass789
4	Nikolay Ivanov	nikolay.vasilyev@example.com	nikolayNewPass234
5	Olga Petrova	olga.newpetrova@example.com	olgaPetrova567
\.


--
-- Name: orders_id_seq; Type: SEQUENCE SET; Schema: public; Owner: store_user
--

SELECT pg_catalog.setval('public.orders_id_seq', 5, true);


--
-- Name: products_id_seq; Type: SEQUENCE SET; Schema: public; Owner: store_user
--

SELECT pg_catalog.setval('public.products_id_seq', 7, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: store_user
--

SELECT pg_catalog.setval('public.users_id_seq', 7, true);


--
-- Name: orderproducts orderproducts_pkey; Type: CONSTRAINT; Schema: public; Owner: store_user
--

ALTER TABLE ONLY public.orderproducts
    ADD CONSTRAINT orderproducts_pkey PRIMARY KEY (order_id, product_id);


--
-- Name: orders orders_pkey; Type: CONSTRAINT; Schema: public; Owner: store_user
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: store_user
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: store_user
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: store_user
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_orderproducts_order_id; Type: INDEX; Schema: public; Owner: store_user
--

CREATE INDEX idx_orderproducts_order_id ON public.orderproducts USING btree (order_id);


--
-- Name: idx_orderproducts_product_id; Type: INDEX; Schema: public; Owner: store_user
--

CREATE INDEX idx_orderproducts_product_id ON public.orderproducts USING btree (product_id);


--
-- Name: idx_orders_order_date; Type: INDEX; Schema: public; Owner: store_user
--

CREATE INDEX idx_orders_order_date ON public.orders USING btree (order_date);


--
-- Name: idx_orders_total_amount; Type: INDEX; Schema: public; Owner: store_user
--

CREATE INDEX idx_orders_total_amount ON public.orders USING btree (total_amount);


--
-- Name: idx_orders_user_id; Type: INDEX; Schema: public; Owner: store_user
--

CREATE INDEX idx_orders_user_id ON public.orders USING btree (user_id);


--
-- Name: idx_products_name; Type: INDEX; Schema: public; Owner: store_user
--

CREATE INDEX idx_products_name ON public.products USING btree (name);


--
-- Name: idx_products_price; Type: INDEX; Schema: public; Owner: store_user
--

CREATE INDEX idx_products_price ON public.products USING btree (price);


--
-- Name: idx_users_email; Type: INDEX; Schema: public; Owner: store_user
--

CREATE INDEX idx_users_email ON public.users USING btree (email);


--
-- Name: idx_users_name; Type: INDEX; Schema: public; Owner: store_user
--

CREATE INDEX idx_users_name ON public.users USING btree (name);


--
-- Name: orderproducts orderproducts_order_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: store_user
--

ALTER TABLE ONLY public.orderproducts
    ADD CONSTRAINT orderproducts_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.orders(id);


--
-- Name: orderproducts orderproducts_product_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: store_user
--

ALTER TABLE ONLY public.orderproducts
    ADD CONSTRAINT orderproducts_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id);


--
-- Name: orders orders_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: store_user
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: pg_database_owner
--

GRANT ALL ON SCHEMA public TO store_user;


--
-- Name: DEFAULT PRIVILEGES FOR TABLES; Type: DEFAULT ACL; Schema: public; Owner: postgres
--

ALTER DEFAULT PRIVILEGES FOR ROLE postgres IN SCHEMA public GRANT ALL ON TABLES TO store_user;


--
-- PostgreSQL database dump complete
--

