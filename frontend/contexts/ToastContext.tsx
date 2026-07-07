"use client";

import React, { createContext, useContext, useState, useCallback, ReactNode, useEffect } from 'react';

export type ToastType = 'success' | 'error' | 'info';

export interface Toast {
  id: string;
  message: string;
  type: ToastType;
}

interface ToastContextType {
  toasts: Toast[];
  addToast: (message: string, type: ToastType) => void;
  removeToast: (id: string) => void;
  error: (message: string) => void;
  success: (message: string) => void;
}

const ToastContext = createContext<ToastContextType | undefined>(undefined);

function ToastItem({ toast, removeToast }: { toast: Toast, removeToast: (id: string) => void }) {
  const [isLeaving, setIsLeaving] = useState(false);
  const [width, setWidth] = useState('100%');

  useEffect(() => {
    const timer = setTimeout(() => {
      setWidth('0%');
    }, 50);

    const leaveTimer = setTimeout(() => {
      setIsLeaving(true);
      setTimeout(() => removeToast(toast.id), 300);
    }, 4000);

    return () => {
      clearTimeout(timer);
      clearTimeout(leaveTimer);
    };
  }, [toast.id, removeToast]);

  const borderColor = toast.type === 'error' ? 'border-red-500' : 
                      toast.type === 'success' ? 'border-emerald-500' : 
                      'border-blue-500';
  
  const bgColor = toast.type === 'error' ? 'bg-red-500' : 
                  toast.type === 'success' ? 'bg-emerald-500' : 
                  'bg-blue-500';

  return (
    <div
      className={`pointer-events-auto relative overflow-hidden flex items-center justify-between p-4 rounded-md shadow-lg bg-slate-900 border-l-4 ${borderColor} ${
        isLeaving ? 'animate-out slide-out-to-bottom-8 fade-out duration-300 fill-mode-forwards' : 'animate-in slide-in-from-bottom-8 fade-in duration-300'
      }`}
    >
      <span className="text-white text-sm font-medium pr-4 relative z-10">{toast.message}</span>
      <button
        onClick={() => {
          setIsLeaving(true);
          setTimeout(() => removeToast(toast.id), 300);
        }}
        className="text-slate-400 hover:text-white transition-colors flex-shrink-0 relative z-10"
        aria-label="Dismiss"
      >
        ✕
      </button>
      <div 
        className={`absolute bottom-0 left-0 h-1 ${bgColor} transition-all ease-linear`}
        style={{ width, transitionDuration: '3950ms' }}
      />
    </div>
  );
}

function ToastContainer({ toasts, removeToast }: { toasts: Toast[], removeToast: (id: string) => void }) {
  if (toasts.length === 0) return null;

  return (
    <div className="fixed bottom-4 right-4 z-50 flex flex-col gap-3 pointer-events-none w-full max-w-sm">
      {toasts.map((toast) => (
        <ToastItem key={toast.id} toast={toast} removeToast={removeToast} />
      ))}
    </div>
  );
}

export function ToastProvider({ children }: { children: ReactNode }) {
  const [toasts, setToasts] = useState<Toast[]>([]);

  const removeToast = useCallback((id: string) => {
    setToasts((prev) => prev.filter((toast) => toast.id !== id));
  }, []);

  const addToast = useCallback((message: string, type: ToastType) => {
    const id = Math.random().toString(36).substring(2, 9);
    setToasts((prev) => [...prev, { id, message, type }]);
  }, []);

  const error = useCallback((message: string) => addToast(message, 'error'), [addToast]);
  const success = useCallback((message: string) => addToast(message, 'success'), [addToast]);

  return (
    <ToastContext.Provider value={{ toasts, addToast, removeToast, error, success }}>
      {children}
      <ToastContainer toasts={toasts} removeToast={removeToast} />
    </ToastContext.Provider>
  );
}

export function useToast() {
  const context = useContext(ToastContext);
  if (!context) {
    throw new Error('useToast must be used within a ToastProvider');
  }
  return context;
}
