import { lightCodeTheme, darkCodeTheme } from 'prism-react-renderer';

const organizationName = '2024-2A-T02-EC11-G01';

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: 'Artemis',
  tagline: 'Restaurando o futuro do planeta agora',
  favicon: 'icons/inteli.png',

  // Set the production url of your site here
  url: `https://${organizationName}.github.io`,
  // Set the /<baseUrl>/ pathname under which your site is served
  // For GitHub pages deployment, it is often '/<projectName>/'
  baseUrl: `/${organizationName}/`,

  projectName: '2024-2A-T02-EC11-G01',
  organizationName: organizationName,
  trailingSlash: false,

  // GitHub pages deployment config.
  // If you aren't using GitHub pages, you don't need these.
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',
  // Even if you don't use internationalization, you can use this field to set
  // useful metadata like html lang. For example, if your site is Chinese, you
  // may want to replace "en" with "zh-Hans".
  i18n: {
    defaultLocale: 'br',
    locales: ['br'],
  },

  presets: [
    [
      'classic',
      /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
      ({
        docs: {
          sidebarPath: './sidebars.ts',
          routeBasePath: '/'
        },
        blog: false,
        theme: {
          customCss: './src/css/custom.css',
        },
      }),
    ],
  ],

  themeConfig:
  /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
  ({
    // Replace with your project's social card
    // image: 'img/backgroud.jpg',
    colorMode: {
      defaultMode: 'dark',
      disableSwitch: false,
      respectPrefersColorScheme: false,
    },
    navbar: {
      title: 'Artemis | Restaurando o futuro do planeta agora',
      logo: {
        alt: 'logo',
        src: 'icons/logo-eng-comp.png',
      },
      items: [
        {
          to: "https://github.com/Inteli-College/2024-2A-T02-EC11-G01",
          position: "right",
          className: "header-github-link",
          "aria-label": "GitHub repository",
        },
      ],
    },
    footer: {
      style: 'dark',
      copyright: `Copyright © ${new Date().getFullYear()} Artemis Inc.`,
    },
    prism: {
      theme: lightCodeTheme,
      darkTheme: darkCodeTheme,
    },
    markdown: {
      mermaid: true,
    },    
  }),
};

export default config;