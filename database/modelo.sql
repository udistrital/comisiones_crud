-- =====================================================
-- ESQUEMA
-- =====================================================

CREATE SCHEMA IF NOT EXISTS comisiones;
ALTER SCHEMA comisiones OWNER TO postgres;
SET search_path TO comisiones;

-- =====================================================
-- TABLAS PARAMÉTRICAS
-- =====================================================

CREATE SEQUENCE tipo_solicitud_id_seq;

CREATE TABLE tipo_solicitud (
    id INTEGER DEFAULT nextval('tipo_solicitud_id_seq'),
    nombre VARCHAR(100),
    descripcion TEXT,
    codigo_abreviacion VARCHAR,
    activo BOOLEAN DEFAULT TRUE,
    fecha_creacion TIMESTAMP WITHOUT TIME ZONE,
    fecha_modificacion TIMESTAMP WITHOUT TIME ZONE,
    CONSTRAINT pk_tipo_solicitud PRIMARY KEY (id)
);

ALTER SEQUENCE tipo_solicitud_id_seq OWNED BY tipo_solicitud.id;


CREATE SEQUENCE estado_solicitud_id_seq;

CREATE TABLE estado_solicitud (
    id INTEGER DEFAULT nextval('estado_solicitud_id_seq'),
    nombre VARCHAR(100),
    descripcion TEXT,
    codigo_abreviacion VARCHAR,
    activo BOOLEAN DEFAULT TRUE,
    fecha_creacion TIMESTAMP WITHOUT TIME ZONE,
    fecha_modificacion TIMESTAMP WITHOUT TIME ZONE,
    CONSTRAINT pk_estado_solicitud PRIMARY KEY (id)
);

ALTER SEQUENCE estado_solicitud_id_seq OWNED BY estado_solicitud.id;


CREATE SEQUENCE tipo_documento_solicitud_id_seq;

CREATE TABLE tipo_documento_solicitud (
    id INTEGER DEFAULT nextval('tipo_documento_solicitud_id_seq'),
    nombre VARCHAR(100),
    descripcion TEXT,
    codigo_abreviacion VARCHAR,
    rol_usuario VARCHAR,
    activo BOOLEAN DEFAULT TRUE,
    fecha_creacion TIMESTAMP WITHOUT TIME ZONE,
    fecha_modificacion TIMESTAMP WITHOUT TIME ZONE,
    CONSTRAINT pk_tipo_documento_solicitud PRIMARY KEY (id)
);

ALTER SEQUENCE tipo_documento_solicitud_id_seq OWNED BY tipo_documento_solicitud.id;


CREATE SEQUENCE tipo_documento_comision_id_seq;

CREATE TABLE tipo_documento_comision (
    id INTEGER DEFAULT nextval('tipo_documento_comision_id_seq'),
    nombre VARCHAR(100),
    descripcion TEXT,
    codigo_abreviacion VARCHAR,
    activo BOOLEAN DEFAULT TRUE,
    fecha_creacion TIMESTAMP WITHOUT TIME ZONE,
    fecha_modificacion TIMESTAMP WITHOUT TIME ZONE,
    CONSTRAINT pk_tipo_documento_comision PRIMARY KEY (id)
);

ALTER SEQUENCE tipo_documento_comision_id_seq OWNED BY tipo_documento_comision.id;


CREATE SEQUENCE estado_comision_id_seq;

CREATE TABLE estado_comision (
    id INTEGER DEFAULT nextval('estado_comision_id_seq'),
    nombre VARCHAR(100),
    descripcion TEXT,
    codigo_abreviacion VARCHAR,
    activo BOOLEAN DEFAULT TRUE,
    fecha_creacion TIMESTAMP WITHOUT TIME ZONE,
    fecha_modificacion TIMESTAMP WITHOUT TIME ZONE,
    CONSTRAINT pk_estado_comision PRIMARY KEY (id)
);

ALTER SEQUENCE estado_comision_id_seq OWNED BY estado_comision.id;


CREATE SEQUENCE tipo_seguimiento_id_seq;

CREATE TABLE tipo_seguimiento (
    id INTEGER DEFAULT nextval('tipo_seguimiento_id_seq'),
    nombre VARCHAR(100),
    descripcion TEXT,
    codigo_abreviacion VARCHAR,
    activo BOOLEAN DEFAULT TRUE,
    fecha_creacion TIMESTAMP WITHOUT TIME ZONE,
    fecha_modificacion TIMESTAMP WITHOUT TIME ZONE,
    CONSTRAINT pk_tipo_seguimiento PRIMARY KEY (id)
);

ALTER SEQUENCE tipo_seguimiento_id_seq OWNED BY tipo_seguimiento.id;


CREATE SEQUENCE estado_documento_id_seq;

CREATE TABLE estado_documento (
    id INTEGER DEFAULT nextval('estado_documento_id_seq'),
    nombre VARCHAR(100),
    descripcion TEXT,
    codigo_abreviacion VARCHAR,
    activo BOOLEAN DEFAULT TRUE,
    fecha_creacion TIMESTAMP WITHOUT TIME ZONE,
    fecha_modificacion TIMESTAMP WITHOUT TIME ZONE,
    CONSTRAINT pk_estado_documento PRIMARY KEY (id)
);

ALTER SEQUENCE estado_documento_id_seq OWNED BY estado_documento.id;


CREATE SEQUENCE estado_documento_comision_id_seq;

CREATE TABLE estado_documento_comision (
    id INTEGER DEFAULT nextval('estado_documento_comision_id_seq'),
    nombre VARCHAR(100),
    descripcion TEXT,
    codigo_abreviacion VARCHAR,
    activo BOOLEAN DEFAULT TRUE,
    fecha_creacion TIMESTAMP WITHOUT TIME ZONE,
    fecha_modificacion TIMESTAMP WITHOUT TIME ZONE,
    CONSTRAINT pk_estado_documento_comision PRIMARY KEY (id)
);

ALTER SEQUENCE estado_documento_comision_id_seq OWNED BY estado_documento_comision.id;

-- =====================================================
-- COMISIÓN
-- =====================================================

CREATE SEQUENCE comision_id_seq;

CREATE TABLE comision (
    id INTEGER DEFAULT nextval('comision_id_seq'),
    descripcion TEXT,
    fecha_inicio TIMESTAMP WITHOUT TIME ZONE,
    fecha_final TIMESTAMP WITHOUT TIME ZONE,
    facultad INTEGER,
    activo BOOLEAN DEFAULT TRUE,
    fecha_creacion TIMESTAMP WITHOUT TIME ZONE,
    fecha_modificacion TIMESTAMP WITHOUT TIME ZONE,
    CONSTRAINT pk_comision PRIMARY KEY (id)
);

ALTER SEQUENCE comision_id_seq OWNED BY comision.id;

-- =====================================================
-- SOLICITUD
-- =====================================================

CREATE SEQUENCE solicitud_id_seq;

CREATE TABLE solicitud (
    id INTEGER DEFAULT nextval('solicitud_id_seq'),
    tercero_id INTEGER,
    tipo_solicitud_id INTEGER,
    comision_id INTEGER,
    observacion_cierre TEXT,
    activo BOOLEAN DEFAULT TRUE,
    fecha_creacion TIMESTAMP WITHOUT TIME ZONE,
    fecha_modificacion TIMESTAMP WITHOUT TIME ZONE,

    CONSTRAINT pk_solicitud PRIMARY KEY (id),

    CONSTRAINT fk_solicitud_tipo_solicitud
        FOREIGN KEY (tipo_solicitud_id)
        REFERENCES tipo_solicitud(id),

    CONSTRAINT fk_solicitud_comision
        FOREIGN KEY (comision_id)
        REFERENCES comision(id)
);

ALTER SEQUENCE solicitud_id_seq OWNED BY solicitud.id;

-- =====================================================
-- INFO FORMULARIO SOLICITUD
-- =====================================================

CREATE SEQUENCE detalle_solicitud_id_seq;

CREATE TABLE detalle_solicitud (
    id INTEGER DEFAULT nextval('detalle_solicitud_id_seq'),
    solicitud_id INTEGER,
    formulario JSONB,
    activo BOOLEAN DEFAULT TRUE,
    fecha_creacion TIMESTAMP WITHOUT TIME ZONE,
    fecha_modificacion TIMESTAMP WITHOUT TIME ZONE,

    CONSTRAINT pk_detalle_solicitud PRIMARY KEY (id),

    CONSTRAINT fk_detalle_solicitud_solicitud
        FOREIGN KEY (solicitud_id)
        REFERENCES solicitud(id)
);

ALTER SEQUENCE detalle_solicitud_id_seq OWNED BY detalle_solicitud.id;

-- =====================================================
-- EVENTOS DE SOLICITUD (FASES)
-- =====================================================

CREATE SEQUENCE historico_estado_solicitud_id_seq;

CREATE TABLE historico_estado_solicitud (
    id INTEGER DEFAULT nextval('historico_estado_solicitud_id_seq'),
    solicitud_id INTEGER,
    estado_solicitud_id INTEGER,
    rol_usuario CHARACTER VARYING,
    tercero_id INTEGER,
    activo BOOLEAN DEFAULT TRUE,
    fecha_creacion TIMESTAMP WITHOUT TIME ZONE,
    fecha_modificacion TIMESTAMP WITHOUT TIME ZONE,

    CONSTRAINT pk_historico_estado_solicitud PRIMARY KEY (id),

    CONSTRAINT fk_historico_estado_solicitud_solicitud
        FOREIGN KEY (solicitud_id)
        REFERENCES solicitud(id),

    CONSTRAINT fk_historico_estado_solicitud_estado_solicitud
        FOREIGN KEY (estado_solicitud_id)
        REFERENCES estado_solicitud(id)
);

ALTER SEQUENCE historico_estado_solicitud_id_seq OWNED BY historico_estado_solicitud.id;

-- =====================================================
-- OBSERVACIONES
-- =====================================================

CREATE SEQUENCE observacion_id_seq;

CREATE TABLE observacion (
    id INTEGER DEFAULT nextval('observacion_id_seq'),
    historico_estado_solicitud_id INTEGER,
    descripcion TEXT,
    activo BOOLEAN DEFAULT TRUE,
    fecha_creacion TIMESTAMP WITHOUT TIME ZONE,
    fecha_modificacion TIMESTAMP WITHOUT TIME ZONE,

    CONSTRAINT pk_observacion PRIMARY KEY (id),

    CONSTRAINT fk_observacion_historico_estado_solicitud
        FOREIGN KEY (historico_estado_solicitud_id)
        REFERENCES historico_estado_solicitud(id)
);

ALTER SEQUENCE observacion_id_seq OWNED BY observacion.id;

-- =====================================================
-- ANEXOS POR FASE DE SOLICITUD
-- =====================================================

CREATE SEQUENCE documento_solicitud_id_seq;

CREATE TABLE documento_solicitud (
    id INTEGER DEFAULT nextval('documento_solicitud_id_seq'),
    documento_id INTEGER,
    historico_estado_solicitud_id INTEGER,
    tipo_documento_id INTEGER,
    estado_documento_id INTEGER,
    activo BOOLEAN DEFAULT TRUE,
    fecha_creacion TIMESTAMP WITHOUT TIME ZONE,
    fecha_modificacion TIMESTAMP WITHOUT TIME ZONE,

    CONSTRAINT pk_documento_solicitud PRIMARY KEY (id),

    CONSTRAINT fk_documento_solicitud_historico_estado_solicitud
        FOREIGN KEY (historico_estado_solicitud_id)
        REFERENCES historico_estado_solicitud(id),

    CONSTRAINT fk_documento_solicitud_tipo_documento_solicitud
        FOREIGN KEY (tipo_documento_id)
        REFERENCES tipo_documento_solicitud(id),

    CONSTRAINT fk_documento_solicitud_estado_documento
        FOREIGN KEY (estado_documento_id)
        REFERENCES estado_documento(id)
);

ALTER SEQUENCE documento_solicitud_id_seq OWNED BY documento_solicitud.id;

-- =====================================================
-- EVENTOS DE COMISIÓN
-- =====================================================

CREATE SEQUENCE historico_estado_comision_id_seq;

CREATE TABLE historico_estado_comision (
    id INTEGER DEFAULT nextval('historico_estado_comision_id_seq'),
    comision_id INTEGER,
    estado_comision_id INTEGER,
    tercero_id INTEGER,
    rol_usuario CHARACTER VARYING,
    descripcion TEXT,
    activo BOOLEAN DEFAULT TRUE,
    fecha_creacion TIMESTAMP WITHOUT TIME ZONE,
    fecha_modificacion TIMESTAMP WITHOUT TIME ZONE,

    CONSTRAINT pk_historico_estado_comision PRIMARY KEY (id),

    CONSTRAINT fk_historico_estado_comision_comision
        FOREIGN KEY (comision_id)
        REFERENCES comision(id),

    CONSTRAINT fk_historico_estado_comision_estado_comision
        FOREIGN KEY (estado_comision_id)
        REFERENCES estado_comision(id)
);

ALTER SEQUENCE historico_estado_comision_id_seq OWNED BY historico_estado_comision.id;

-- =====================================================
-- ANEXOS POR EVENTO DE COMISIÓN
-- =====================================================

CREATE SEQUENCE documento_comision_id_seq;

CREATE TABLE documento_comision (
    id INTEGER DEFAULT nextval('documento_comision_id_seq'),
    documento_id INTEGER,
    historico_estado_comision_id INTEGER,
    tipo_documento_id INTEGER,
    estado_documento_comision_id INTEGER,
    activo BOOLEAN DEFAULT TRUE,
    fecha_creacion TIMESTAMP WITHOUT TIME ZONE,
    fecha_modificacion TIMESTAMP WITHOUT TIME ZONE,

    CONSTRAINT pk_documento_comision PRIMARY KEY (id),

    CONSTRAINT fk_documento_comision_historico_estado_comision
        FOREIGN KEY (historico_estado_comision_id)
        REFERENCES historico_estado_comision(id),

    CONSTRAINT fk_documento_comision_tipo_documento_comision
        FOREIGN KEY (tipo_documento_id)
        REFERENCES tipo_documento_comision(id),

    CONSTRAINT fk_documento_comision_estado_documento_comision
        FOREIGN KEY (estado_documento_comision_id)
        REFERENCES estado_documento_comision(id)
);

ALTER SEQUENCE documento_comision_id_seq OWNED BY documento_comision.id;

-- =====================================================
-- SEGUIMIENTO
-- =====================================================

CREATE SEQUENCE seguimiento_id_seq;

CREATE TABLE seguimiento (
    id INTEGER DEFAULT nextval('seguimiento_id_seq'),
    historico_estado_comision_id INTEGER,
    tipo_seguimiento_id INTEGER,
    descripcion TEXT,
    activo BOOLEAN DEFAULT TRUE,
    fecha_creacion TIMESTAMP WITHOUT TIME ZONE,
    fecha_modificacion TIMESTAMP WITHOUT TIME ZONE,

    CONSTRAINT pk_seguimiento PRIMARY KEY (id),

    CONSTRAINT fk_seguimiento_historico_estado_comision
        FOREIGN KEY (historico_estado_comision_id)
        REFERENCES historico_estado_comision(id),

    CONSTRAINT fk_seguimiento_tipo_seguimiento
        FOREIGN KEY (tipo_seguimiento_id)
        REFERENCES tipo_seguimiento(id)
);

ALTER SEQUENCE seguimiento_id_seq OWNED BY seguimiento.id;
