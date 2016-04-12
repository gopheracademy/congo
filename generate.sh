#! /bin/bash

goagen --design github.com/gopheracademy/congo/design app
goagen --design github.com/gopheracademy/congo/design client 
goagen --design github.com/gopheracademy/congo/design js 
goagen --design github.com/gopheracademy/congo/design schema
goagen --design github.com/gopheracademy/congo/design swagger
goagen --design github.com/gopheracademy/congo/design gen --pkg-path github.com/goadesign/gorma

