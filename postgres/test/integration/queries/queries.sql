@query GetExamples:Rows
    %%SelectAllFromExample%%
@endquery

@query GetExampleById:Row
    %%SelectAllFromExample%%
    WHERE "id" = $string:id
@endquery

@query CreateExample:Rows
    INSERT INTO "example" ("id", "name")
    VALUES ($string:id, $string:name)
@endquery

@query UpdateExample:Rows
    UPDATE "example"
    SET "name" = $string:name
    WHERE "id" = $string:id
@endquery

@query DeleteExample:Rows
    DELETE FROM "example"
    WHERE "id" = $string:id
@endquery
