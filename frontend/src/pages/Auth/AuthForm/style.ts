import {css} from 'emotion'

import {FlexVert, FlexHor} from "../../../common/styles/positioning";
import {Font312Red} from "../../../common/styles/fonts";
import colors from "../../../common/styles/colors";

export const style = {
    FlexVert,
    FlexHor,
    Font312Red,
    Form: css`
        width: 500px;
    
        padding: 16px;
        
        margin: 250px auto auto;
        border: 1px solid ${colors.greenPrimary};
        box-sizing: border-box;
        border-radius: 4px;
        box-shadow: 0 0 10px 2px rgba(0, 0, 0, 0.17);
  `,
    ButtonsBlock: css`
        display: flex;
        justify-content: flex-end;
        margin-top: 8px;
    `,
    ButtonsBlockWithErrors: css`
        margin-top: 8px;
        justify-content: space-between;
    `,
    AuthButton: css`
        min-width: 100px;
        height: 40px;
        margin-left: 8px;
    `,
    ValidationBlock: css`
        margin: 0;
        padding-left: 12px;
    `,
    ButtonsWrapper: css`
      display: flex;
    `
};