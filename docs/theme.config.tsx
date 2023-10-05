/* eslint-disable import/no-anonymous-default-export */
export default {
    project: {
      link: 'https://github.com/nexentra/aesir'
    },
    docsRepositoryBase:"https://github.com/nexentra/aesir",
    logo: <strong>Project</strong>,
    sidebar:{
      toggleButton:true,
      autoCollapse:true,
    },
    toc:{
      float:true,
      backToTop:true,
    },
    primaryHue: { dark: 180, light: 257 },
    primarySaturation: { dark: 80, light: 80 },
    navigation:{
      prev:true,
      next:true,
    },
    gitTimestamp: () => {
      return (
        <div className="hidden">
          {/* <span>Updated at {timestamp.toLocaleString()}</span> */}
        </div>
      )
    },
    editLink:{
      component:() => {
        return (
          <div className="hidden">
            {/* <span>Updated at {timestamp.toLocaleString()}</span> */}
          </div>
        )
      },
    },
    feedback:{
      
    },
      footer: {
        text: (
          <span>
            Nexentra {new Date().getFullYear()} ©{' '}
            <a href="https://nextra.site" target="_blank">
              Aesir
            </a>
            .
          </span>
        )
      },

  }