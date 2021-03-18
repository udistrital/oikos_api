ALTER TABLE oikos.espacio_fisico ADD column tipo_edificacion_id INTEGER;
ALTER TABLE oikos.espacio_fisico ADD column tipo_terreno_id INTEGER;
COMMENT ON COLUMN oikos.espacio_fisico.tipo_edificacion_id IS 'Hace referencia al id al tipo de edificación en el api parametros_crud';
COMMENT ON COLUMN oikos.espacio_fisico.tipo_terreno_id IS 'Hace referencia al id al tipo de terreno en el api parametros_crud';
