import {css} from 'emotion'

import {
    Font312DarkGray, Font312RedAttention, Font312Black
} from "../../styles/fonts";
import {FlexVert, FlexHorCenter} from "../../styles/positioning";
import colors from "../../styles/colors";

const maxHeight = 400;

export const style = {
    Font312DarkGray,
    Font312RedAttention,
    Font312Black,
    FlexVert,
    FlexHorCenter,
    maxHeight,

    Dropdown: css`
      position: absolute;
      z-index: 2;
      top: 56px;
      max-height: ${maxHeight}px;
      width: 100%;
      background-color: ${colors.white};
      box-shadow: rgba(0, 0, 0, 0.16) 0px 4px 10px 0px;
      border-right: 1px solid ${colors.greenActive};
      border-left: 1px solid ${colors.greenActive};
      border-bottom: 1px solid ${colors.greenActive};
      border-bottom-left-radius: 4px;
      border-bottom-right-radius: 4px;
      box-sizing: border-box;
    `,

    SingleLine: css`
      cursor: pointer;
      padding: 4px 8px;
      height: 30px;
      border-bottom: 1px solid ${colors.darkGray};
      
      &:hover {
        background-color: ${colors.greenLight};
      }
      
      &:last-of-type {
        border: 0;
      }
    `,

    SelectedLine: css`
      background-color: ${colors.greenLight};
    `
}