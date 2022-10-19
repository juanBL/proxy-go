Creacion de un proxy que añada headers a la URL destino.

Tipo: curl -x localhost:8000 http://httpbin.org/anything

El username del proxy será la APIKEY del usuario.
El password serán los headers a enviar a la URL destino.

Adicionalmente el proxy tendrá que:
- Comprobar que la apikey existe en MySQL y es válida (no ha expirado). La tabla tiene 2 fields: apikey (uuid) + expiration date.
- Loggear en una lista Redis la config del request, en formato Json (LPUSH logged-requests <json>). El JSON a loggear puede tener esta forma, por ejemplo:
{
"url": "http://httpbin.org/anything",
"apikey": "484afe2c-140f-46b5-88f9-170db60d94bd",
"headers": [{"header1": "value1"}, {"header2": "value2"}]
}

El log anterior sería generado por el siguiente comando:

curl -x http://484afe2c-140f-46b5-88f9-170db60d94bd:header1=value1&header2=value2@localhost:8000 http://httpbin.org/anything

Donde:
    apikey = 484afe2c-140f-46b5-88f9-170db60d94bd
    proxy = localhost:8000
    headers = header1=value1 + header2=value2
    url destino = http://httpbin.org/anything

Si el proxy funciona bien, deberías ver los headers enviados desde el proxy en la URL destino (http://httpbin.org/anything)