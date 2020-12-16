import { ResponseError, Context } from 'umi-request';
import {
  NavBarProps,
  TitleListItem,
  NavBarListItem,
  TabBarProps,
  TabBarListItem,
  history
} from 'alita';

import HomeGary from './assets/demoIcon/home.png';
import HomeBlue from './assets/demoIcon/home1.png';
import ListGary from './assets/demoIcon/list.png';
import ListBlue from './assets/demoIcon/list1.png';
import SetGary from './assets/demoIcon/setting.png';
import SetBlue from './assets/demoIcon/setting1.png';

// 请求中间件 就是发起请求和响应之后需要统一操作数据就写这
// https://github.com/umijs/umi-request#example-1
const middleware = async (ctx: Context, next: any) => {
  console.log('a1');
  await next();
  console.log('a2');
};

export const request = {
  prefix: '', // 统一的请求头
  middlewares: [middleware],
  errorHandler: (error: ResponseError) => {
    // 集中处理错误
    console.log(error);
  },
};


const titleList: TitleListItem[] = [
  {
    pagePath: '/',
    title: '首页',
  },
  {
    pagePath: '/list',
    title: '列表',
  },
  {
    pagePath: '/settings',
    title: '设置',
  },
  {
    pagePath: '/noteFolder',
    title: '笔记文件夹',
  },
  {
    pagePath: '/noteDetails',
    title: '笔记详情',
  },
  {
    pagePath: '/noteList',
    title: '笔记列表',
  },
  {
    pagePath: '/createNote',
    title: '创建笔记',
  },
  {
    pagePath: '/admin',
    title: '管理',
  },
];
const navList: NavBarListItem[] = [
  {
    pagePath: '/login',
    navBar: {
      hideNavBar: true,
    }
  },
];
const navBar: NavBarProps = {
  navList,
  fixed: true,
  onLeftClick: () => {
    // router.goBack();
  },
};
const tabList: TabBarListItem[] = [
  {
    pagePath: '/',
    text: '首页',
    iconPath: HomeGary,
    selectedIconPath: HomeBlue,
    title: '首页',
    iconSize: '',
    badge: '',
  },
  {
    pagePath: '/noteFolder',
    text: '文件夹',
    iconPath: SetGary,
    selectedIconPath: SetBlue,
    title: '笔记文件夹',
    iconSize: '',
    badge: '',
  },
  {
    pagePath: '/settings',
    text: '设置',
    iconPath: SetGary,
    selectedIconPath: SetBlue,
    title: '设置',
    iconSize: '',
    badge: '',
  },
];

const tabBar: TabBarProps = {
  color: `#999999`,
  selectedColor: '#00A0FF',
  borderStyle: 'white',
  position: 'bottom',
  list: tabList,
};

export const mobileLayout = {
  documentTitle: '默认标题',
  navBar,
  tabBar,
  titleList,
};

// if(localStorage.getItem('userInfo')) {
//   history.replace('/')
// } else {
//   history.replace('/login')
// }