# AuthService
A service to providing the authentication feature such as:

- Register flow
- Facebook Oatuh2 flow
- Google Oauth2 flow

Additionally, this server also contains the coupon management and email sending feature.

This server is a server providing the feature handler and gateway. Here are the APIs:

<table>
<tr>
  <th>API NAME</th>
  <th>URL</th>
  <th>Method</th>
  <th>Description</th>
  <th>Request</th>
  <th>Response</th>
</tr>
<tr>
  <td>Healthcheck</td>
  <td>/ping</td>
  <td>GET</td>
  <td>Healthcheck endpoint. Response message when server alive.</td>
  <td>N/A</td>
  <td>200: 
  <code>{"message": "pong"}`</code></td>
</tr>
<tr>
  <td>Register</td>
  <td>/register</td>
  <td>POST</td>
  <td>Register a new user, if user already register than response the JWT token only</td>
  <td>
    <code>
    {
      "username": "Lucas",
      "email": "chucobo5219@gmail.com",
      "password": "klasdfjvier123"
    }
    </code>
  </td>
  <td>
  200:
  <code>
    {
        "Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6Ikx1Y2FzIiwiZXhwIjoxNTg3NTcxNDYzfQ.KxKRNYmS6Q-07JZNIpNYNyjXJ9p1l6bacBAcp94zPDc"
    }
  </code><br>
  502:
    <code>
    {
      "error": "Username is require in request body"
    }
    </code><br>
  502:
    <code>
    {
      "error": "the password should contains at lease 8 charactor"
    }
    </code><br>
  502:
    <code>
    {
      "error": "incorrect email ex: 1234@domain.com current input:chucobo5219gmail.com"
    }
    </code><br>
  </td>
  </tr>
  <tr>
    <td>Get google oauth url</td>
    <td>/get_googleoauth_url</td>
    <td>GET</td>
    <td>Get google oauth login endpoint</td>
    <td>N/A</td>
    <td>200: 
    <code>
    {
    "url": "https://accounts.google.com/o/oauth2/auth?client_id=253768931865-or5h985g6ftkljaeqd8obveq2eod4a2i.apps.googleusercontent.com&redirect_uri=https%3A%2F%2Famazingtalker.herokuapp.com%2Fgoogleoauth&response_type=code&scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.profile&state=state"
}
    </code></td>
  </tr>
  <tr>
    <td>Get facebook oauth url</td>
    <td>/get_facebookoauth_url</td>
    <td>GET</td>
    <td>Get facebook oauth login endpoint</td>
    <td>N/A</td>
    <td>200: 
    <code>
{
    "url": "https://www.facebook.com/v3.2/dialog/oauth?client_id=3043148599098072&redirect_uri=https%3A%2F%2Famazingtalker.herokuapp.com%2Ffacebookoauth&response_type=code&scope=public_profile+email&state=state"
}
    </code></td>
</table>

## Oatuh
To passing the oauth, we need to get the oauth url and the Oauth2 process is passing if the token can get the user profile. The user profile is getting from oauth provider server. When profile is invalid. The error message is responded.

## How to test
The server is using heroku:
https://amazingtalker.herokuapp.com/

You can using above above domain to end to end testing each endpoints.

For example: https://amazingtalker.herokuapp.com/ping to testing the healthcheck endpoint.