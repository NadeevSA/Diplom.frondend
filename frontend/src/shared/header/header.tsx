import React, { useEffect } from 'react';
import { Header , HeaderLogo, HeaderModule } from '@consta/uikit/Header';
import authServer from '../../serviceAuth/authServer';
import { User } from '@consta/uikit/User';
import { ModalLogin } from '../../modals/login/login';
import { SideBar } from './sideBar';

export const CustomHeader = () => {
  const [UserName, setUserName] = React.useState<string>("");
  const [UserEmail, setUserEmail] = React.useState<string>("");

  function isLogin(login: string){
      debugger;
      setUserName(login);
  }

  useEffect(() => {
    debugger;
    Avatar();
  }, [UserName, UserEmail]);

  function Avatar() {
    return (
      <div>
        {
          authServer.getToken() != "" ?
          <User
            name={UserName}
            info={UserEmail}
            view="ghost"
            size="l"
          /> :
          <ModalLogin userName={setUserName} userEmail={setUserEmail}></ModalLogin>
        }
      </div> 
    );
  };

  return (
    <Header
  leftSide={
    <>
    <HeaderModule>
        <SideBar userName={UserName} isLogin={isLogin}></SideBar>
    </HeaderModule>
    <HeaderModule indent="m">
      <HeaderLogo>
        <img width={150} height={50} src="./logoAppRunner.png" alt="Логотип"/>
      </HeaderLogo>
    </HeaderModule>
    </>
  }
  rightSide={
    <>
      <HeaderModule indent="s">
          <Avatar></Avatar>
        </HeaderModule>
      <HeaderModule indent="s">  
      </HeaderModule>
    </>
  }
/>)}
