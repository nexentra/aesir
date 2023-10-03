// @ts-nocheck
import { useEffect } from 'react';
import { useRouter } from 'next/router';
import dynamic from 'next/dynamic';
import '../globals.css';
import Head from 'next/head';
import Script from 'next/script';

import { Toaster } from '@/shadcn/components/ui/toaster';
import AnimatedBackground from '@/components/ui/animated-background';
import { ThemeProvider } from '@/components/ui/theme-provider';
import CustomCursor from '@/components/ui/cursor';
import NextTopLoader from 'nextjs-toploader';

function App({ Component, pageProps }) {
    const router = useRouter();

    useEffect(() => {
        window.scrollTo(0, 0);
        const handleRouteChange = (url) => {
            window.gtag('config', 'UA-1200538-30', { page_path: url, path_url: url });
        };
        router.events.on('routeChangeComplete', handleRouteChange);
        return () => router.events.off('routeChangeComplete', handleRouteChange);
    }, [router.events]);

    return (
        <>
            <Head>
                <meta
                    name="viewport"
                    content="minimum-scale=1, initial-scale=1, width=device-width, shrink-to-fit=no, user-scalable=no, viewport-fit=cover"
                />
            </Head>
            <NextTopLoader />
            <AnimatedBackground />
            <ThemeProvider
                attribute="class"
                defaultTheme="system"
                enableSystem
            >
                <CustomCursor />
                <Component {...pageProps} />
            </ThemeProvider>
            <Toaster />


            <Script
                strategy="afterInteractive"
                src="https://www.googletagmanager.com/gtag/js?id=UA-1200538-30"
            />
            <Script id="google-analytics" strategy="afterInteractive">
                {`
          window.dataLayer = window.dataLayer || [];
          function gtag(){window.dataLayer.push(arguments);}
          gtag('js', new Date());
          gtag('config', 'UA-1200538-30', {
            page_path: window.location.pathname,
          });
          window.gtag = gtag;
        `}
            </Script>
        </>
    );
}

export default App;