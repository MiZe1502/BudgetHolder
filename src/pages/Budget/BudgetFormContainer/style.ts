import {css} from "emotion";

import {Font312Black} from "../../../common/styles/fonts";
import {FlexVert, SideMainPadding} from "../../../common/styles/positioning";

import colors from "../../../common/styles/colors";

export const style = {
    Font312Black,

    FlexVert,
    SideMainPadding,

    Wrapper: css`
      width: 100%;
      min-height: 150px;
      border: 1px solid ${colors.darkGray};
      border-radius: 4px;
      display: flex;
    `,

    Container: css`
      margin-bottom: 40px;
    `
}
