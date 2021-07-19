#Bookings and Reservations

This is repository for my bookings and reservations project.

- Built in Go version 1.16
- Uses the [chi router](https://github.com/go-chi/chi)
- Uses [alex edwards SCS session management](https://github.com/alexedwards/scs)
- Uses [nosurf](https://github.com/justinas/nosurf)

To remember:

- Migration generation table:
soda generate fizz *CreateNameTable*

- Migration generate FK/Index:
soda generate fizz *CreateFKNameTable*

- Seed generation:
soda generate sql *SeedNameTable*