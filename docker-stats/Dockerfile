FROM docker:17.12.0
RUN apk update && apk add bash
ADD run.sh /run.sh
RUN chmod +x /run.sh
ENTRYPOINT ["/run.sh"]
