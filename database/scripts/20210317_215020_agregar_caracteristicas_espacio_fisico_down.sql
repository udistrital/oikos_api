ALTER TABLE oikos.espacio_fisico DROP column tipo_edificacion_id INTEGER;
ALTER TABLE oikos.espacio_fisico DROP column tipo_terreno_id INTEGER;
COMMENT ON COLUMN oikos.espacio_fisico.tipo_edificacion_id IS '';
COMMENT ON COLUMN oikos.espacio_fisico.tipo_terreno_id IS '';
