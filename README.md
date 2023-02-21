# totp-client
A proof of concept project to write a cli client to generate TOTP codes. Not for production use. This is an academic exercise to understand the TOTP protocols.

# HOTP

HOTP (HMAC-based One-Time Password) was published as an informational IETF [RFC 4226](https://datatracker.ietf.org/doc/html/rfc4226#ref-BCK1). HMAC stands for Hashed Message Authentication Code.

The HTOP algorithm is defined as the following where:

K = shared secret between client and server; each HOTP
           generator has a different and unique secret K

C = 8-byte counter value, the moving factor.  This counter
           MUST be synchronized between the HOTP generator (client)
           and the HOTP validator (server)

The `Truncate` function represents the function that converts an HMAC-SHA-1
       value into an HOTP value

```
HOTP(K,C) = Truncate(HMAC-SHA-1(K,C))
```

Steps are as follows:

1. Generate an HMAC-SHA-1 value. `Let HS = HMAC-SHA-1(K,C)`. This is 20-byte, binary string
1. Use Dynamic Truncation (DT) to convert the 20-byte string into a 31-bit string. `Let Sbits = DT(HS)`
1. Compute an HOTP value. `Let Snum  = StToNum(Sbits); Return D = Snum mod 10^Digit`

Steps 2 and 3 above are part of the DT function. The purpose of
   the dynamic offset truncation technique is to extract a 4-byte
   dynamic binary code from a 160-bit (20-byte) HMAC-SHA-1 result.

HTOP must produce at least 6 digit code and possibly a 7 or 8 digits. All my codes use 8 digits. 



# How totp works
Time-based One Time Passwords are typically defined by [RFC 6238](https://datatracker.ietf.org/doc/html/rfc6238)

TOTP is an extension of HOTP (HMAC-based One-Time Password).
Hashed Message Authentication Code (HMAC)

Both TOTP and HOTP were written by the Initiative for Open Authentication (OATH). An industry group dedicated to creating open protocols for authentication.

`TOTP = HOTP(K, T)`
K = Key
T = (Current Unix time - T0) / X, where the
   default floor function is used in the computation.

