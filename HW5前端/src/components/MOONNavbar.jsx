import React from 'react';
import Container from 'react-bootstrap/Container';
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';

/**
 * Renders the navbar component with a sign-in or sign-out button depending on whether or not a user is authenticated
 * @param props
 */
export const MOONNavbar = (props) => {
    return (
        <>
            <Navbar className="navbarStyle">
            <Container>
                <Navbar.Brand href="/">
                    <img
                    src="/img/home.png"
                    width="60"
                    height="60"
                    className="d-inline-block align-center"
                    alt="React Bootstrap logo"
                    />{' '}
                    電腦零組件販售屋
                </Navbar.Brand>
                
                <Navbar.Toggle aria-controls="basic-navbar-nav" />
                <Navbar.Collapse id="basic-navbar-nav">
                <Nav className="me-auto">
                    <Nav.Link href="/products">產品</Nav.Link>
                    <Nav.Link href="/customers">客戶</Nav.Link>
                    <Nav.Link href="/orders">訂單</Nav.Link>
                    <Nav.Link href="/items">貨物</Nav.Link>
                </Nav>
                </Navbar.Collapse>
                </Container>
            </Navbar>
        </>
    );   
};
