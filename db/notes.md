

```sql
CREATE DATABASE scenariobank
  ENCODING 'UTF8';

\c scenariobank;
 ```

  ```sql
CREATE TABLE scenarioer (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created TIMESTAMP NOT NULL DEFAULT NOW(),
    expires TIMESTAMP NOT NULL
);

CREATE INDEX idx_scenarioer_created ON scenarioer(created);

\d+ scenarioer


ALTER TABLE public.scenarioer
ADD COLUMN scenario_type TEXT NOT NULL DEFAULT 'unknown';

```

```sql
INSERT INTO scenarioer (title, content, created, expires) VALUES (
    'Manglende tilgangsstyring i fagsystem',
    'Et sentralt fagsystem mangler tilstrekkelig rolle- og tilgangsstyring.\n
Ansatte kan få tilgang til informasjon utover tjenstlig behov, uten at dette
fanges opp av eksisterende logger eller kontroller.\n
\nDette kan føre til brudd på personvernregelverket, tap av tillit og
omdømmeskade dersom uautorisert innsyn avdekkes.',
    NOW(),
    NOW() + INTERVAL '365 days'
);

INSERT INTO scenarioer (title, content, created, expires) VALUES (
    'Avhengighet til ekstern skyleverandør uten exit-plan',
    'Virksomheten er sterkt avhengig av én ekstern skyleverandør for kritiske
tjenester.\n
Det finnes ingen dokumentert eller testet exit-strategi dersom leverandøren
endrer vilkår, får langvarig driftsstans eller blir utilgjengelig av
geopolitiske årsaker.\n
\nScenarioet kan gi betydelige konsekvenser for tilgjengelighet og
tjenesteleveranse.',
    NOW(),
    NOW() + INTERVAL '365 days'
);


INSERT INTO scenarioer (title, content, created, expires) VALUES (
    'Mangelfull logging ved sikkerhetshendelser',
    'Ved en mulig sikkerhetshendelse oppdages det at logger er ufullstendige
eller mangler sentrale hendelser.\n
Dette gjør det vanskelig å gjennomføre effektiv hendelseshåndtering,
etterforskning og varsling innenfor lovpålagte frister.\n
\nScenarioet øker risikoen for feil håndtering av hendelser og manglende
etterlevelse av regelverk.',
    NOW(),
    NOW() + INTERVAL '7 days'
);


SELECT id, title, created, expires
FROM scenarioer
ORDER BY created;

```

```sql
CREATE ROLE web
  LOGIN
  PASSWORD 'pass';

GRANT CONNECT ON DATABASE scenariobank TO web;

GRANT SELECT, INSERT, UPDATE, DELETE
ON ALL TABLES IN SCHEMA public
TO web;

ALTER DEFAULT PRIVILEGES
IN SCHEMA public
GRANT SELECT, INSERT, UPDATE, DELETE
ON TABLES
TO web;

GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO web;

ALTER DEFAULT PRIVILEGES
IN SCHEMA public
GRANT USAGE, SELECT ON SEQUENCES
TO web;
```

```sh
psql -h localhost -p 5432 -U web scenariobank
```