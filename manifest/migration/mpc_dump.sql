--
-- PostgreSQL database dump
--

-- Dumped from database version 12.3 (Debian 12.3-1.pgdg100+1)
-- Dumped by pg_dump version 12.16 (Ubuntu 12.16-0ubuntu0.20.04.1)

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
-- Name: mpc_context; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mpc_context (
    user_id character varying NOT NULL,
    context character varying,
    updated_at timestamp(6) without time zone,
    request character varying,
    token character varying,
    created_at timestamp(6) without time zone,
    deleted_at timestamp(6) without time zone,
    pub_key character varying,
    token_data character varying
);


ALTER TABLE public.mpc_context OWNER TO postgres;

--
-- Name: wallet_addr; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wallet_addr (
    user_id character varying NOT NULL,
    wallet_addr character varying(255) NOT NULL,
    chain_id bigint NOT NULL
);


ALTER TABLE public.wallet_addr OWNER TO postgres;

--
-- Name: mpc_context mpc_context_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mpc_context
    ADD CONSTRAINT mpc_context_pkey PRIMARY KEY (user_id);


--
-- Name: wallet_addr_user_id_wallet_addr_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX wallet_addr_user_id_wallet_addr_idx ON public.wallet_addr USING btree (user_id, wallet_addr);


--
-- Name: wallet_addr_wallet_addr_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX wallet_addr_wallet_addr_idx ON public.wallet_addr USING btree (wallet_addr);


--
-- PostgreSQL database dump complete
--

