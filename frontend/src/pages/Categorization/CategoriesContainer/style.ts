import {css} from 'emotion'

import {SideMainPadding} from "../../../common/styles/positioning";
import colors from "../../../common/styles/colors";

const maxHeight = 500;

export const style = {
    SideMainPadding,
    maxHeight,

    Container: css`
      border: 1px solid ${colors.darkGray};
      border-radius: 4px;
      width: 100%;
      max-height: ${maxHeight}px;
      min-height: 250px;
      overflow: scroll;
      display: flex;
    `,

    ContainerWithoutData: css`
      justify-content: center;
      align-items: center;
    `
}
