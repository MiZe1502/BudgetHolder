import {css} from 'emotion'

export {
    SideMinorPadding, FlexVert, FlexHorCenter
} from "../../../../common/styles/positioning";
export {Font716Black, Font312Black} from "../../../../common/styles/fonts";

import colors from "../../../../common/styles/colors";

const imageHeight = 44;

export const Image = css`
    height: ${imageHeight}px;
    
    &:hover {
      cursor: pointer;
    }
`;

export const UserBlock = css`
    background-color: ${colors.greenPrimary};
    border-radius: 50%;
    position: relative;
`

export const Login = css`
    margin-right: 8px;
`;

//TODO: create common popup dropdown container

export const Popup = css`
    width: 200px;
    position: absolute;
    top: ${imageHeight}px;
    right: 0px;
    background-color: ${colors.white};
    z-index: 2;
    box-shadow: rgba(0, 0, 0, 0.16) 0px 4px 10px 0px;
    border: 1px solid ${colors.greenActive};
    border-radius: 4px;
    box-sizing: border-box;
`;

export const PopupItem = css`
    cursor: pointer;
    padding: 4px 8px;
    height: 30px;
    background-color: ${colors.white};
    border-bottom: 1px solid ${colors.darkGray};

    &:hover {
      cursor: pointer;
      background-color: ${colors.greenLight};
    }
    
    &:first-of-type {
      border-top-left-radius: 4px;
      border-top-right-radius: 4px;
    }
    
    &:last-of-type {
      border: 0;
      border-bottom-left-radius: 4px;
      border-bottom-right-radius: 4px;
    }
`;