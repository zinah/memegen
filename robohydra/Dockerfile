FROM node:16

RUN mkdir /test-server

WORKDIR /test-server

RUN npm install robohydra

# TODO investigate why this is necessary?
RUN npm install -g robohydra

RUN mkdir /test-server/robohydra

RUN mkdir /test-server/assets

COPY plugins /test-server/robohydra/plugins

COPY test.conf /test-server/robohydra/test.conf

COPY assets/default_test_cat.png /test-server/assets/default_test_cat.png

CMD ["robohydra", "robohydra/test.conf"]
