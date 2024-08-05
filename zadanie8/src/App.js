import React, { useState } from 'react';
import axios from 'axios';
import './App.css';

const App = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [message, setMessage] = useState('');
  const [token, setToken] = useState('');

  const handleRegister = async () => {
    try {
      const response = await axios.post('http://localhost:8080/register', { username, password });
      console.log("Registration response:", response);
      setMessage(response.data.message);
    } catch (error) {
      console.error("Registration error:", error);
      setMessage(error.response?.data?.error || 'Registration failed');
    }
  };

  const handleLogin = async () => {
    try {
      const response = await axios.post('http://localhost:8080/login', { username, password });
      console.log("Login response:", response);
      setToken(response.data.token);
      setMessage('Login successful');
    } catch (error) {
      console.error("Login error:", error);
      setMessage(error.response?.data?.error || 'Login failed');
    }
  };

  const fetchHello = async () => {
    try {
      const response = await axios.get('http://localhost:8080/auth/hello', {
        headers: { Authorization: token }
      });
      console.log("Fetch hello response:", response);
      setMessage(response.data.message);
    } catch (error) {
      console.error("Fetch hello error:", error);
      setMessage('Could not fetch hello message');
    }
  };

  return (
    <div className="App">
      <h1>Login/Registration</h1>
      <input
        type="text"
        placeholder="Username"
        value={username}
        onChange={(e) => setUsername(e.target.value)}
      />
      <input
        type="password"
        placeholder="Password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
      />
      <button onClick={handleRegister}>Register</button>
      <button onClick={handleLogin}>Login</button>
      {token && <button onClick={fetchHello}>Fetch Hello</button>}
      {message && <p>{message}</p>}
    </div>
  );
};

export default App;

