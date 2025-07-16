// Este contexto de React gestiona el estado de autenticación del usuario, incluyendo login, logout y el estado de completado del perfil, persistiendo los datos en el almacenamiento local.

import React, { createContext, useState, useContext, useEffect } from 'react';

const AuthContext = createContext(null);

export const AuthProvider = ({ children }) => {
  const [user, setUser] = useState(null);

  useEffect(() => {
    const storedUser = localStorage.getItem('user');
    if (storedUser) {
      setUser(JSON.parse(storedUser));
    }
  }, []);

  const login = (userData) => {
    localStorage.setItem('authToken', userData.token);
    localStorage.setItem('user', JSON.stringify(userData));
    setUser({
      isAuthenticated: true,
      ...userData,
    });
  };

  const logout = () => {
    localStorage.removeItem('authToken');
    localStorage.removeItem('user');
    setUser(null);
  };

  const completeProfile = () => {
    setUser((currentUser) => {
      const updatedUser = { ...currentUser, profileCompleted: true };
      localStorage.setItem('user', JSON.stringify(updatedUser));
      return updatedUser;
    });
  };

  return (
    <AuthContext.Provider value={{ user, login, logout, completeProfile }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => useContext(AuthContext);