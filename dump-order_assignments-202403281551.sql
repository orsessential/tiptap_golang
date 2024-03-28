PGDMP  ;    3                |            order_assignments     16.2 (Ubuntu 16.2-1.pgdg20.04+1)     16.2 (Ubuntu 16.2-1.pgdg20.04+1)                0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                       0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false                       1262    16388    order_assignments    DATABASE     }   CREATE DATABASE order_assignments WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.UTF-8';
 !   DROP DATABASE order_assignments;
                postgres    false                        2615    2200    public    SCHEMA        CREATE SCHEMA public;
    DROP SCHEMA public;
                pg_database_owner    false                       0    0    SCHEMA public    COMMENT     6   COMMENT ON SCHEMA public IS 'standard public schema';
                   pg_database_owner    false    4            �            1259    16449    items    TABLE     8  CREATE TABLE public.items (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    description text,
    quantity integer NOT NULL,
    order_id integer,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);
    DROP TABLE public.items;
       public         heap    postgres    false    4            �            1259    16448    items_id_seq    SEQUENCE     �   CREATE SEQUENCE public.items_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.items_id_seq;
       public          postgres    false    218    4                       0    0    items_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.items_id_seq OWNED BY public.items.id;
          public          postgres    false    217            �            1259    16439    orders    TABLE     =  CREATE TABLE public.orders (
    id integer NOT NULL,
    customer_name character varying(255) NOT NULL,
    ordered_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);
    DROP TABLE public.orders;
       public         heap    postgres    false    4            �            1259    16438    orders_id_seq    SEQUENCE     �   CREATE SEQUENCE public.orders_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 $   DROP SEQUENCE public.orders_id_seq;
       public          postgres    false    216    4                       0    0    orders_id_seq    SEQUENCE OWNED BY     ?   ALTER SEQUENCE public.orders_id_seq OWNED BY public.orders.id;
          public          postgres    false    215            z           2604    16452    items id    DEFAULT     d   ALTER TABLE ONLY public.items ALTER COLUMN id SET DEFAULT nextval('public.items_id_seq'::regclass);
 7   ALTER TABLE public.items ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    217    218    218            v           2604    16442 	   orders id    DEFAULT     f   ALTER TABLE ONLY public.orders ALTER COLUMN id SET DEFAULT nextval('public.orders_id_seq'::regclass);
 8   ALTER TABLE public.orders ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    216    215    216                      0    16449    items 
   TABLE DATA           b   COPY public.items (id, name, description, quantity, order_id, created_at, updated_at) FROM stdin;
    public          postgres    false    218   `                 0    16439    orders 
   TABLE DATA           W   COPY public.orders (id, customer_name, ordered_at, created_at, updated_at) FROM stdin;
    public          postgres    false    216   �                  0    0    items_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.items_id_seq', 5, true);
          public          postgres    false    217                       0    0    orders_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.orders_id_seq', 3, true);
          public          postgres    false    215            �           2606    16458    items items_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.items DROP CONSTRAINT items_pkey;
       public            postgres    false    218            ~           2606    16447    orders orders_pkey 
   CONSTRAINT     P   ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);
 <   ALTER TABLE ONLY public.orders DROP CONSTRAINT orders_pkey;
       public            postgres    false    216            �           2606    16459    items items_order_id_fkey    FK CONSTRAINT     z   ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.orders(id);
 C   ALTER TABLE ONLY public.items DROP CONSTRAINT items_order_id_fkey;
       public          postgres    false    216    218    3198               c   x�3�����K54�2��9�8�9��Lt�u�,M������LMM,�Hq��LJ�tM�,A7��� h8�~cC+cs=SC#3<R\1z\\\ ��'P         c   x�}�!�0u��@��ki���SK� � ������0G�?��lZJR��\�h�U���D�:�s������,�k��Q�^z�`�7ٳ��e�!�     