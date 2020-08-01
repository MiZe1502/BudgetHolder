import {css} from "emotion";

export {Font312Black} from "../../../common/styles/fonts";
export {FlexVert, SideMainPadding} from "../../../common/styles/positioning";

import colors from "../../../common/styles/colors";

export const Wrapper = css`
  width: 100%;
  min-height: 150px;
  border: 1px solid ${colors.darkGray};
  border-radius: 4px;
  padding: 20px;
`;

export const Container = css`
  margin-bottom: 40px;
`;