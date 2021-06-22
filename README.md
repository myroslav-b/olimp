### Про проєкт

Проєкт реалізує REST проксі до **Реєстру суб'єктів освітньої діяльності ЄДЕБО** ([registry.edbo.gov.ua](https://registry.edbo.gov.ua))

#### Обґрунтування

Реєстр ЄДЕБО пропонує REST API для доступу до даних. Однак, форма запиту і формат отримуваних даних відрізняються для [закладів загальної середньої освіти](https://registry.edbo.gov.ua/opendata/institutions/) та [інших категорій навчальних закладів](https://registry.edbo.gov.ua/opendata/universities/) (вищої, фахової передвищої
та професійної освіти).

Проєкт:
* уніфікує формат запиту і результатів, що повертаються;
* кешує дані отримані з первинного реєстру;
* архітектурно дозволяє легко підключати інші (додаткові) реєстри та змінювати формат даних, що повертаються.

#### Методи API

##### Довідники

**_GET /api/v1/catalog/{name}_**

_name_ - довідник

* _fi_ - перелік полів, що повертаються запитом до реєстру;
* _it_ - коди категорій навчальних закладів (установ) для яких можна отримати дані;
* _rc_ - коди регіонів (областей, міст обласного підпорядкування);
* _tr_ - дозволені типи запитів.

##### Отримання даних з реєстру

**_GET /api/v1/register/{inst}/{reg}_**

* _inst_ - код категорії навчальних закладів (установ);
* _reg_ - код регіону.

Для уточнення запиту використовуються два види параметрів:

* _r.{fi}_ - значенням параметра є один із дозволених типів запитів (_tr_);
* _v.{fi}_ - значення параметра конкретизує умову запиту.

Як правило, параметри _r.{fi}_ та _v.{fi}_ використовуються парами, однак запити типу "empty/notempty" параметра _v.{fi}_ не потребують.

Наприклад:

_/api/v1/register/3/26_

Повертає перелік (масив JSON об'єктів) закладів загальної середньої освіти Івано-Франківської області.

_/api/v1/register/3/26?r.website=notempty_

Повертає перелік (масив JSON об'єктів) закладів загальної середньої освіти Івано-Франківської області для яких вказані адреси веб-сайтів.

_/api/v1/register/3/26?r.website=notempty&r.name=contains&v.name=ліцей_

Повертає перелік (масив JSON об'єктів) закладів загальної середньої освіти Івано-Франківської області назви яких містять рядок "ліцей" і для яких вказані адреси веб-сайтів.

Якщо у "парних" параметрах відсутній один із елементів, то такі параметри будуть проігноровані. Наприклад, запити

_/api/v1/register/3/26?r.website=notempty&r.name=contains_

або

_/api/v1/register/3/26?r.website=notempty&v.name=ліцей_

повернуть перелік закладів загальної середньої освіти Івано-Франківської області для яких вказані адреси веб-сайтів.

#### Параметри запуску

* short:"a" long:"address" env:"OLIMP_ADDRESS" default:":8080" description:"server address";
* short:"s" long:"source" env:"OLIMP_SOURCE_REGISTRY" choice:"edbo" default:"edbo" description:"source of institution registry";
* short:"t" long:"templife" env:"OLIMP_TEMP_LIFE" default:"86400" description:"lifetime of batches".

#### Docker

**_docker-compose.yml_**

```yaml
version: '3.1'

services:
  olimp:
    image: docker.pkg.github.com/myroslav-b/olimp/olimp:latest
    restart: always
    networks:
      - "traefik_default"
    labels:
      - "traefik.enable=true"
      - "traefik.docker.network=traefik_default"
      - "traefik.http.routers.olimp-router.rule=Host(`olimp.oippo.if.ua`)"
      - "traefik.http.routers.olimp-router.tls=true"
      - "traefik.http.routers.olimp-router.tls.certresolver=letsEncrypt"
      - "traefik.http.routers.olimp-router.entrypoints=websecure"
      - "traefik.http.routers.olimp-router.service=olimp-web-srv"
      - "traefik.http.services.olimp-web-srv.loadbalancer.server.port=8080"
networks:
  traefik_default:
    external: true
```

**_traefik.yml_**

```yaml
entryPoints:
  web:
    address: ":80"
    http:
      redirections:
        entrypoint:
          to: websecure
          scheme: https

  websecure:
    address: ":443"

providers:
  docker:
    network: "traefik_default"
    endpoint: "unix:///var/run/docker.sock"
    exposedByDefault: false

certificatesResolvers:
  letsEncrypt:
    acme:
      email: "mymail@mail.net"
      storage: "acme.json"
      httpChallenge:
        entryPoint: web

api:
  dashboard: true
```
