// Este componente de React renderiza la página de inicio de la aplicación, incluyendo la barra de navegación, una sección principal destacada, características del servicio, una galería de productos y el pie de página.

import React from 'react';
import { Link } from 'react-router-dom';
import { FaClipboardList, FaTruck, FaShieldAlt } from 'react-icons/fa';

import '../Styles/home.css';

const Home = () => {
  return (
    <div className="home-page-container">
      <header className="home-navbar">
        <div className="navbar-container">
          <Link to="/" className="navbar-logo">Sistema de Gas</Link>
          <nav>
            <Link to="/login" className="nav-link">Iniciar Sesión</Link>
            <Link to="/register" className="nav-button">Registrarse</Link>
          </nav>
        </div>
      </header>

      <main>
        <section className="hero-section">
          <div className="hero-container">
            <div className="hero-content fade-in-up">
              <h1>La gestión de tu gas, <br /> más fácil que nunca.</h1>
              <p>Pide, rastrea y gestiona tus bombonas de gas desde la comodidad de tu hogar o negocio. Simple, rápido y seguro.</p>
              <div className="hero-buttons">
                <Link to="/register" className="btn-primary">Empezar Ahora</Link>
                <Link to="/login" className="btn-secondary">Ya tengo cuenta</Link>
              </div>
            </div>
            <div className="hero-image">
              <img src="/ilustracion2.jpg" alt="Bombonas de gas en un almacén" />
            </div>
          </div>
        </section>

        <section className="features-section">
          <div className="section-container">
            <h2>Todo lo que necesitas, en un solo lugar</h2>
            <div className="features-grid">
              <div className="feature-card reveal">
                <FaClipboardList className="feature-icon" />
                <h3>Pedidos Simplificados</h3>
                <p>Crea y gestiona tus pedidos en pocos clics, sin llamadas ni complicaciones.</p>
              </div>
              <div className="feature-card reveal">
                <FaTruck className="feature-icon" />
                <h3>Seguimiento de Entregas</h3>
                <p>Conoce el estado de tus entregas en tiempo real y prepárate para recibirlas.</p>
              </div>
              <div className="feature-card reveal">
                <FaShieldAlt className="feature-icon" />
                <h3>Gestión Segura</h3>
                <p>Tu información y tus pedidos están protegidos en nuestra plataforma confiable.</p>
              </div>
            </div>
          </div>
        </section>

        <section className="products-section">
            <div className="section-container">
                <h2>Nuestros Productos</h2>
                <div className="products-grid">
                    <div className="product-card reveal">
                        <img src="/ilustracion01.png" alt="Bombona de 10 Kg" className="product-image"/>
                        <h3>Bombona Social (10 Kg)</h3>
                        <p>Ideal para el hogar y familias pequeñas.</p>
                    </div>
                     <div className="product-card reveal">
                        <img src="/ilustracion02.png" alt="Bombona de 18 Kg" className="product-image commercial"/>
                        <h3>Bombona Comercial (18 Kg)</h3>
                        <p>Perfecta para pequeños negocios y restaurantes.</p>
                    </div>
                     <div className="product-card reveal">
                        <img src="/ilustracion03.png" alt="Bombona de 43 Kg" className="product-image industrial"/>
                        <h3>Bombona Industrial (43 Kg)</h3>
                        <p>La solución para grandes consumos y cocinas industriales.</p>
                    </div>
                </div>
            </div>
        </section>
      </main>

      <footer className="home-footer">
        <p>© {new Date().getFullYear()} Sistema de Gas. Todos los derechos reservados.</p>
      </footer>
    </div>
  );
};

export default Home;