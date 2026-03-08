import { NextResponse } from 'next/server'
import type { NextRequest } from 'next/server'

export const proxy = (request: NextRequest) => {
    const userRole = request.cookies.get('username')?.value
    const isProtectedRoute = request.nextUrl.pathname.startsWith('/dashboard')
    const isAuthRoute = request.nextUrl.pathname === '/login' || request.nextUrl.pathname === '/signup'
    if (isProtectedRoute && !userRole) {
        return NextResponse.redirect(new URL('/login', request.url))
    }
    if (isAuthRoute && userRole) {
        return NextResponse.redirect(new URL('/dashboard', request.url))
    }
    return NextResponse.next()
}

export const config = {
    matcher: ['/dashboard/:path*', '/login', '/signup'],
}