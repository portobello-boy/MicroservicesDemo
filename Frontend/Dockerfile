FROM node

COPY package.json package-lock.json ./

RUN npm i

RUN npm i --save @angular/cli

COPY . .

RUN ng serve