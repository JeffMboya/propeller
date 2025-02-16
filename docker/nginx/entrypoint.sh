#!/bin/ash

if [ -z "$MG_MQTT_CLUSTER" ]
then
      envsubst '${MG_MQTT_ADAPTER_MQTT_PORT}' < /etc/nginx/snippets/mqtt-upstream-single.conf > /etc/nginx/snippets/mqtt-upstream.conf
      envsubst '${MG_MQTT_ADAPTER_WS_PORT}' < /etc/nginx/snippets/mqtt-ws-upstream-single.conf > /etc/nginx/snippets/mqtt-ws-upstream.conf
else
      envsubst '${MG_MQTT_ADAPTER_MQTT_PORT}' < /etc/nginx/snippets/mqtt-upstream-cluster.conf > /etc/nginx/snippets/mqtt-upstream.conf
      envsubst '${MG_MQTT_ADAPTER_WS_PORT}' < /etc/nginx/snippets/mqtt-ws-upstream-cluster.conf > /etc/nginx/snippets/mqtt-ws-upstream.conf
fi

envsubst '
    ${MG_NGINX_SERVER_NAME}
    ${MG_AUTH_HTTP_PORT}
    ${MG_USERS_HTTP_PORT}
    ${MG_THINGS_HTTP_PORT}
    ${MG_THINGS_AUTH_HTTP_PORT}
    ${MG_NGINX_MQTT_PORT}
    ${MG_NGINX_MQTTS_PORT}' < /etc/nginx/nginx.conf.template > /etc/nginx/nginx.conf

exec nginx -g "daemon off;"
