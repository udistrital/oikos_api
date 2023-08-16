CREATE TABLE IF NOT EXISTS oikos.centro_costos (
    id SERIAL NOT NULL,
    dependencia_id INTEGER NOT NULL,
    sede_id INTEGER NOT NULL,
    codigo VARCHAR(12) NOT NULL,
    nombre TEXT NOT NULL,
    activo BOOLEAN NOT NULL,
    fecha_creacion TIMESTAMP NOT NULL,
    fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT uq_codigo_centro_costos UNIQUE (codigo),
    CONSTRAINT pk_centro_costos PRIMARY KEY (id)
);

ALTER TABLE oikos.centro_costos
    ADD CONSTRAINT fk_centro_costos_dependencia_id FOREIGN KEY (dependencia_id)
    REFERENCES oikos.dependencia (id) MATCH FULL
    ON DELETE RESTRICT ON UPDATE CASCADE;

ALTER TABLE oikos.centro_costos
    ADD CONSTRAINT fk_centro_costos_sede_id FOREIGN KEY (sede_id)
    REFERENCES oikos.espacio_fisico (id) MATCH FULL
    ON DELETE RESTRICT ON UPDATE CASCADE;
