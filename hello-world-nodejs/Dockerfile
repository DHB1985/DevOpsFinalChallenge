#Se cambia la imagen del contenedor para que sea más liviana.
FROM node:14-alpine 

WORKDIR /app 

COPY package*.json ./ 

RUN npm install 

COPY . . 

#Se expone el puerto 3000 para poder utilizarlo cuando se levanta el contenedor.
EXPOSE 3000

CMD ["npm", "run", "start"]
