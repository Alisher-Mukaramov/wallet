--
-- PostgreSQL database dump
--

-- Dumped from database version 14.2 (Debian 14.2-1.pgdg110+1)
-- Dumped by pg_dump version 14.1

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

--
-- Name: confirmer(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.confirmer() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN

     IF NEW.state_id = 2 and NEW.processed_date is not null THEN
      
      UPDATE accounts SET saldo = saldo + NEW.amount
      WHERE id = NEW.debit_acc_id AND currency_id = NEW.currency_id;
      
      UPDATE accounts SET saldo = saldo - NEW.amount
      WHERE id = NEW.credit_acc_id AND currency_id = NEW.currency_id;
      
      
     END IF;
     
    RETURN NULL;
END
$$;


ALTER FUNCTION public.confirmer() OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: account_types; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.account_types (
    id bigint NOT NULL,
    name text NOT NULL
);


ALTER TABLE public.account_types OWNER TO postgres;

--
-- Name: account_types_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.account_types ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.account_types_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: accounts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.accounts (
    id bigint NOT NULL,
    account character varying(20),
    user_id bigint,
    currency_id bigint,
    saldo numeric(20,2),
    comment text,
    acc_type_id bigint NOT NULL
);


ALTER TABLE public.accounts OWNER TO postgres;

--
-- Name: accounts_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.accounts ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.accounts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: clients; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.clients (
    id bigint NOT NULL,
    name text,
    second_name text,
    patronymic text,
    date_of_birth date,
    country text,
    city text,
    postal_addr text,
    residential_addr text,
    passport_series text,
    passport_issued text,
    passport_exp_date date
);


ALTER TABLE public.clients OWNER TO postgres;

--
-- Name: clients_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.clients ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.clients_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: currencies; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.currencies (
    id bigint NOT NULL,
    name text NOT NULL,
    code character varying(3) NOT NULL
);


ALTER TABLE public.currencies OWNER TO postgres;

--
-- Name: currencies_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.currencies ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.currencies_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: limits; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.limits (
    id bigint NOT NULL,
    max_balance numeric(20,2),
    trn_per_day numeric(10,0),
    trn_per_month numeric(10,0),
    user_type_id bigint
);


ALTER TABLE public.limits OWNER TO postgres;

--
-- Name: limits_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.limits ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.limits_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: states; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.states (
    id bigint NOT NULL,
    code character varying(20) NOT NULL
);


ALTER TABLE public.states OWNER TO postgres;

--
-- Name: states_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.states ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.states_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: transaction_types; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transaction_types (
    id bigint NOT NULL,
    name text NOT NULL
);


ALTER TABLE public.transaction_types OWNER TO postgres;

--
-- Name: transaction_types_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.transaction_types ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.transaction_types_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: transactions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transactions (
    id bigint NOT NULL,
    insert_date date NOT NULL,
    processed_date date,
    amount numeric(20,2) NOT NULL,
    debit_acc_id bigint NOT NULL,
    credit_acc_id bigint NOT NULL,
    currency_id bigint NOT NULL,
    trn_type_id bigint NOT NULL,
    state_id bigint NOT NULL
);


ALTER TABLE public.transactions OWNER TO postgres;

--
-- Name: transactions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.transactions ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.transactions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    email text NOT NULL,
    phone text NOT NULL,
    is_active boolean NOT NULL,
    client_id bigint,
    user_type_id bigint NOT NULL,
    secret_key character varying(15)
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.users ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: user_types; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_types (
    id bigint NOT NULL,
    code text
);


ALTER TABLE public.user_types OWNER TO postgres;

--
-- Name: user_types_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.user_types ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.user_types_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Data for Name: account_types; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.account_types (id, name) FROM stdin;
1	Пользовательский
2	Технический
3	Технически-комиссионный
\.


--
-- Data for Name: accounts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.accounts (id, account, user_id, currency_id, saldo, comment, acc_type_id) FROM stdin;
3	20216870000000000000	1	1	0.00	USD account	1
1	20216000000000000000	1	2	96263.70	TJS account	1
4	20216000000000000000	2	2	7.50	TJS account	1
2	10120000000000000000	\N	2	-71218.00	Технический TJS счет для пополнения/списания	2
\.


--
-- Data for Name: clients; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.clients (id, name, second_name, patronymic, date_of_birth, country, city, postal_addr, residential_addr, passport_series, passport_issued, passport_exp_date) FROM stdin;
1	Алексейхон	Турдихуджаев	Богайратович	2000-10-10	Точикистон	Восеъ	734000	Кадучи	223344332	10-10-2016	2026-10-10
\.


--
-- Data for Name: currencies; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.currencies (id, name, code) FROM stdin;
1	Доллар	USD
2	Сомони	TJS
3	Рубль	RUB
\.


--
-- Data for Name: limits; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.limits (id, max_balance, trn_per_day, trn_per_month, user_type_id) FROM stdin;
1	10000.00	20	100	2
2	100000.00	40	200	1
\.


--
-- Data for Name: states; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.states (id, code) FROM stdin;
1	PENDING
2	SUCCESS
3	ERROR
\.


--
-- Data for Name: transaction_types; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.transaction_types (id, name) FROM stdin;
1	Пополнение
2	Списание
\.


--
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.transactions (id, insert_date, processed_date, amount, debit_acc_id, credit_acc_id, currency_id, trn_type_id, state_id) FROM stdin;
14	2022-07-11	2022-07-11	13000.00	1	2	2	1	2
15	2022-07-11	2022-07-12	13000.00	1	2	2	1	2
16	2022-07-11	2022-07-12	13000.00	1	2	2	1	2
19	2022-07-12	2022-07-12	13000.00	1	2	2	1	2
20	2022-07-12	2022-07-12	13000.00	1	2	2	1	2
21	2022-07-12	2022-07-12	3.00	1	2	2	1	2
22	2022-07-12	2022-07-12	7.50	1	2	2	1	2
23	2022-07-12	2022-07-12	7.50	4	2	2	1	2
\.


--
-- Data for Name: user_types; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.user_types (id, code) FROM stdin;
1	Не идентифицирован
2	Идентифицирован
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, email, phone, is_active, client_id, user_type_id, secret_key) FROM stdin;
1	bilol.mukaramov@icloud.com	+992900900900	t	\N	1	$eCrEt
2	alekseykhon@gmail.com	+992918918918	t	1	2	@lek$eykhon
\.


--
-- Name: account_types_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.account_types_id_seq', 3, true);


--
-- Name: accounts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.accounts_id_seq', 4, true);


--
-- Name: clients_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.clients_id_seq', 1, true);


--
-- Name: currencies_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.currencies_id_seq', 3, true);


--
-- Name: limits_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.limits_id_seq', 2, true);


--
-- Name: states_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.states_id_seq', 3, true);


--
-- Name: transaction_types_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.transaction_types_id_seq', 2, true);


--
-- Name: transactions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.transactions_id_seq', 23, true);


--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_id_seq', 2, true);


--
-- Name: user_types_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_types_id_seq', 2, true);


--
-- Name: account_types account_types_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.account_types
    ADD CONSTRAINT account_types_pkey PRIMARY KEY (id);


--
-- Name: accounts accounts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_pkey PRIMARY KEY (id);


--
-- Name: clients clients_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.clients
    ADD CONSTRAINT clients_pkey PRIMARY KEY (id);


--
-- Name: currencies currencies_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.currencies
    ADD CONSTRAINT currencies_pkey PRIMARY KEY (id);


--
-- Name: states states_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.states
    ADD CONSTRAINT states_pkey PRIMARY KEY (id);


--
-- Name: transaction_types transaction_types_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transaction_types
    ADD CONSTRAINT transaction_types_pkey PRIMARY KEY (id);


--
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- Name: users user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- Name: user_types user_types_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_types
    ADD CONSTRAINT user_types_pkey PRIMARY KEY (id);


--
-- Name: fki_clients_fki; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fki_clients_fki ON public.users USING btree (client_id);


--
-- Name: transactions confirmer; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER confirmer AFTER UPDATE ON public.transactions FOR EACH ROW EXECUTE FUNCTION public.confirmer();


--
-- Name: accounts account_type_fki; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT account_type_fki FOREIGN KEY (acc_type_id) REFERENCES public.account_types(id) NOT VALID;


--
-- Name: users clients_fki; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT clients_fki FOREIGN KEY (client_id) REFERENCES public.clients(id) NOT VALID;


--
-- Name: transactions cred_acc_fki; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT cred_acc_fki FOREIGN KEY (credit_acc_id) REFERENCES public.accounts(id) NOT VALID;


--
-- Name: accounts currency_fki; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT currency_fki FOREIGN KEY (currency_id) REFERENCES public.currencies(id);


--
-- Name: transactions currency_fki; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT currency_fki FOREIGN KEY (currency_id) REFERENCES public.currencies(id) NOT VALID;


--
-- Name: transactions debt_acc_fki; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT debt_acc_fki FOREIGN KEY (debit_acc_id) REFERENCES public.accounts(id) NOT VALID;


--
-- Name: transactions state_fki; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT state_fki FOREIGN KEY (state_id) REFERENCES public.states(id) NOT VALID;


--
-- Name: transactions trn_type_fki; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT trn_type_fki FOREIGN KEY (trn_type_id) REFERENCES public.transaction_types(id) NOT VALID;


--
-- Name: accounts user_fki; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT user_fki FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: limits user_type_fki; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.limits
    ADD CONSTRAINT user_type_fki FOREIGN KEY (user_type_id) REFERENCES public.user_types(id);


--
-- Name: users user_types_fki; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT user_types_fki FOREIGN KEY (user_type_id) REFERENCES public.user_types(id) NOT VALID;


--
-- PostgreSQL database dump complete
--

