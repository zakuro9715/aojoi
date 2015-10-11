FROM heroku/go:1.4.2
MAINTAINER Takahiro Yaota <zakuro@yuzakuro.me>

RUN cp /app/user/src/github.com/zakuro9715/aojoi/templates /app/user/ -R
RUN cp /app/user/src/github.com/zakuro9715/aojoi/static /app/user/ -R

CMD ["aojoi"]
