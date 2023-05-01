import React from 'react';
import './HomePage.css';
import logo from './logo.png'
import {Link} from "react-router-dom";


function NavBar() {
    return (
        <nav>
            <div id={"navbarlogo"}>
                <img src={logo} alt="Logo" />
            </div>
            <ul>
                <li>
                    <Link to="/login" id="loginButton">Login</Link>
                </li>
                <li>
                    <Link to="/register" id="registerButton">Register</Link>
                </li>
            </ul>
        </nav>
    );
}

function Content() {
    return (
        <main>
            <h1>Welcome to My Messenger App!</h1>
            <p>Get started by logging in or registering for an account.</p>
        </main>
    );
}

function Footer() {
    return (
        <footer>
            <p>&copy; 2023 My Messenger App</p>
        </footer>
    );
}

function HomePage() {
    return (
        <div>
            <NavBar />
            <Content />
            <Footer />
        </div>
    );
}

export default HomePage;