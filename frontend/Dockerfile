FROM node:14.3

WORKDIR /app

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

RUN apt-get update -y && \
  apt-get upgrade -y
#   npm init -y && \
#   npm install npm@6.14.4 && \
#   npm install ejs@3.0.1 &&\
#   npm install defu@3.1.0

ADD . /app

RUN npm run build

EXPOSE 3001

CMD ["npm", "run", "start"]